package e2e

import (
	"context"
	"database/sql"
)

type TodoService interface {
	AddTodo(ctx context.Context, title, description string) (*Todo, error)
	GetTodos(ctx context.Context) ([]*Todo, error)
	ToggleTodo(ctx context.Context, id int) (*Todo, error)
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *service {
	return &service{db: db}
}

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

func (s *service) GetTodos(ctx context.Context) ([]*Todo, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, description, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		t := &Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (s *service) GetTodo(ctx context.Context, id int) (*Todo, error) {
	row := s.db.QueryRowContext(ctx, "SELECT id, title, description, done FROM todos WHERE id = ?", id)

	t := &Todo{}
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Done)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *service) ToggleTodo(ctx context.Context, id int) (*Todo, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE todos SET done = CASE done WHEN 0 THEN 1 ELSE 0 END WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return s.GetTodo(ctx, id)
}
