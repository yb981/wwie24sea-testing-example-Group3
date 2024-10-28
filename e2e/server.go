package e2e

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type server struct {
	db          *sql.DB
	todoService TodoService
}

func (s *server) Start(ctx context.Context, addr string) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "CREATE TABLE todos (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, description TEXT, done BOOLEAN DEFAULT FALSE)")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO todos (title, description) VALUES ('test', 'test')")
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	mux := http.NewServeMux()
	mux.Handle("POST /todos", ContentTypeJSON(s.AddTodo))
	mux.Handle("GET /todos", ContentTypeJSON(s.GetTodos))
	mux.Handle("PATCH /todos/{id}/toggle", ContentTypeJSON(s.ToggleTodo))

	return http.ListenAndServe(addr, mux)
}

func NewServer(db *sql.DB) *server {
	service := NewService(db)
	return &server{db: db, todoService: service}
}

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (s *server) AddTodo(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := s.todoService.AddTodo(r.Context(), todo.Title, todo.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.todoService.GetTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	numericID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	t, err := s.todoService.ToggleTodo(r.Context(), numericID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
