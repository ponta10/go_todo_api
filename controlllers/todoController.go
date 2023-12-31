package controlllers

import (
    "encoding/json"
    "net/http"
    "go_todo_api/models"
)

type TodoController struct {
    Model *models.TodoModel
}

// TodoControllerを返す新しいインスタンス生成し初期設定を行う
// ポインタ型を使うことで同じインスタンスを共有し(同じものの共有)、その状態を変更することができます
func NewTodoController(m *models.TodoModel) *TodoController {
    return &TodoController{Model: m}
}

// レシーバーは、メソッドが属している型を指定するためのもので、そのメソッドを呼び出すために必要な情報を提供します
func (h *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // h(TodoaHandlerのポインタ)のフィールドModelにアクセス
    // 同じインスタンスを異なるメソッド間で共有したり、インスタンスの状態を変更したりすることができます。
    // ModelのAll関数を呼び出す
    todos, err := h.Model.All()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // todosをJSON形式で返す
    json.NewEncoder(w).Encode(todos)
}

func (h *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
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
