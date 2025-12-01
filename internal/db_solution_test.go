package internal

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSolutionDB(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define expectations
	// 1. Expect query for name
	mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test User"))
}

func TestSolutionDB_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expect error on first query
	mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	if err := SolutionDB(db); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
