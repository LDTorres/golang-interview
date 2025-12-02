package http

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	repo UserRepository
}

func (u *UserController) CreateEvent(c *fiber.Ctx) error {
	var event User
	if err := c.BodyParser(&event); err != nil {
		return err
	}

	if event.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	if len(event.Title) > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title must be less than 100 characters",
		})
	}

	if event.StartTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Start time is required",
		})
	}

	if event.EndTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "End time is required",
		})
	}

	parsedStartTime, err := time.Parse(time.RFC3339, event.StartTime)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid start time format",
		})
	}

	parsedEndTime, err := time.Parse(time.RFC3339, event.EndTime)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid end time format",
		})
	}

	if parsedStartTime.After(parsedEndTime) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Start time must be before end time",
		})
	}

	event.CreatedAt = time.Now().Format(time.RFC3339)

	id, err := u.repo.CreateEvent(&event)

	if err != nil {
		return err
	}

	event.ID = fmt.Sprintf("%d", id)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"event": event,
	})
}

func (u *UserController) ListEvents(c *fiber.Ctx) error {
	events, err := u.repo.ListEvents()

	if err != nil {
		return err
	}

	if events == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No events found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})
}

func (u *UserController) GetEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	event, err := u.repo.GetEvent(id)

	if err != nil {
		return err
	}

	if event == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}
