package models

import (
    "database/sql"
)

type Todo struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
}

// データベースとの接続
type TodoModel struct {
    DB *sql.DB
}

//この設計により、TodoModel のインスタンスはデータベースへの接続やクエリの実行などの操作に利用することができます
func NewTodoModel(DB *sql.DB) *TodoModel {
    return &TodoModel{DB: DB}
}

// レシーバーは、メソッドが属している型を指定するためのもので、そのメソッドを呼び出すために必要な情報を提供します
func (m *TodoModel) All() ([]Todo, error) {
    rows, err := m.DB.Query("SELECT id, task FROM todos")
    if err != nil {
        return nil, err
    }
    // DBの取得・保持はメモリを使うため、closeする
    defer rows.Close()

    var todos []Todo
    // rowsが続く限りスープ(JSのforEach)
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }

    return todos, nil
}

func (m *TodoModel) Insert(task string) (int, error) {
    result, err := m.DB.Exec("INSERT INTO todos (task) VALUES (?)", task)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}
