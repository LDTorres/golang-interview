package http

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateEvent(event *User) (id int64, err error)
	ListEvents() ([]User, error)
	GetEvent(id string) (*User, error)
}

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		conn: db,
	}
}

func (u *userRepository) CreateEvent(event *User) (id int64, err error) {
	event.ID = uuid.New().String()

	_, err = u.conn.Exec("INSERT INTO events (id, title, description, start_time, end_time, created_at) VALUES ($1, $2, $3, $4, $5, $6)", event.ID, event.Title, event.Description, event.StartTime, event.EndTime, event.CreatedAt)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (u *userRepository) ListEvents() ([]User, error) {
	rows, err := u.conn.Query("SELECT * FROM events ORDER BY start_time ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []User

	for rows.Next() {
		var event User

		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.StartTime, &event.EndTime, &event.CreatedAt); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (u *userRepository) GetEvent(id string) (*User, error) {
	user := &User{}

	err := u.conn.QueryRow("SELECT * FROM events WHERE id = $1", id).Scan(&user.ID, &user.Title, &user.Description, &user.StartTime, &user.EndTime, &user.CreatedAt)

	return user, err
}
