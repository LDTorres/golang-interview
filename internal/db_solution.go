package internal

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PublicID     string    `json:"public_id"`
	MessageCount int       `json:"message_count"`
}

func SolutionDB(db *sql.DB) error {
	// Query by name
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = $1", 1).Scan(&name)
	if err != nil {
		return err
	}

	fmt.Println("SelectedName", name)

	// Query all object
	var users []User
	rows, err := db.Query("SELECT id, created_at, updated_at, name, email, public_id, message_count FROM users")
	if err != nil {
		fmt.Println("Err", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Email, &user.PublicID, &user.MessageCount); err != nil {
			fmt.Println("Err", err)
			return err
		}
		users = append(users, user)
	}

	fmt.Printf("SelectedUsers %v", users)

	return nil
}
