package unittests_test

//go test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mocking for the database (we don't need real database access here)
type mockDB struct {
	sql.DB
}

// Assume this is the actual function we are testing
func NewService(db *sql.DB) *service {
	return &service{db: db}
}

// Dummy service structure
type service struct {
	db *sql.DB
}

// Test for NewService function
func TestNewService(t *testing.T) {
	// 1. Create a mock database object
	mockDatabase := &mockDB{}

	// 2. Call the NewService function and store the result
	s := NewService(&mockDatabase.DB)

	// 3. Check if the returned service object is non-nil
	assert.NotNil(t, s, "Expected service to be non-nil")

	// 4. Verify that the db field in the service structure is set correctly
	assert.Equal(t, &mockDatabase.DB, s.db, "Expected the db field to be set correctly")
}

// Todo structure (this must be present in your code)
type Todo struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

// Method to add a todo item
func (s *service) AddTodo(ctx context.Context, title, description string) (*Todo, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO todos (title, description) VALUES (?, ?)", title, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:          int(id),
		Title:       title,
		Description: description,
		Done:        false,
	}, nil
}

// Test for AddTodo function
func TestAddTodo(t *testing.T) {
	// 1. Create a sqlmock object
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// 2. Create an instance of the service with the mock database
	s := NewService(db)

	// 3. Define input values for the test
	ctx := context.Background()
	title := "Test Todo"
	description := "This is a test todo item."

	// 4. Set up the expected mock for the database call
	mock.ExpectExec("INSERT INTO todos").
		WithArgs(title, description).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulates LastInsertId() = 1 and 1 affected row

	// 5. Call the AddTodo function and check the result
	todo, err := s.AddTodo(ctx, title, description)

	// 6. Verify that no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, todo)

	// 7. Check the properties of the returned Todo object
	assert.Equal(t, 1, todo.ID, "Expected the ID to be 1")
	assert.Equal(t, title, todo.Title, "Expected the title to match")
	assert.Equal(t, description, todo.Description, "Expected the description to match")
	assert.False(t, todo.Done, "Expected Done to be false")

	// 8. Verify that all expected SQL commands were called
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
