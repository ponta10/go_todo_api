package handlers

import (
    "encoding/json"
    "net/http"
    "go_todo_api/models" // <- Replace with your project's path
)

type TodoHandler struct {
    Model *models.TodoModel
}

func NewTodoHandler(m *models.TodoModel) *TodoHandler {
    return &TodoHandler{Model: m}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    todos, err := h.Model.All()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var todo models.Todo
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := h.Model.Insert(todo.Task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    todo.ID = id
    json.NewEncoder(w).Encode(todo)
}
