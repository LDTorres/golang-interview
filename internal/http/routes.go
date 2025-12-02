package http

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	repo := &userRepository{conn: db}
	controller := &UserController{repo: repo}

	app.Get("/events", controller.ListEvents)
	app.Post("/events", controller.CreateEvent)
	app.Get("/events/:id", controller.GetEvent)
}
