package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "go_todo_api/handler"
    "go_todo_api/models"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // dbの設定
    db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // DBとの対話を担当するモデルのインスタンス作成
    todoModel := models.NewTodoModel(db)
    // HTTPリクエストの処理
    // 引数にtodoModelを渡し、処理内でDBとの対話を依頼
    // ここでtodoHandler.goでの記述のおかげで同じ初期設定を使いまわせる
    // そしてtodohandler内の関数を使うことができる
    todoHandler := handlers.NewTodoHandler(todoModel)

    router := mux.NewRouter()
    // GET /todosの処理がきたらtodoHandlerのGetTodosを呼び出す
    // インスタンス.関数で呼び出せる
    router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
    // POST /todosの処理がきたらtodoHandlerのCreateTodosを呼び出す
    router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

    fmt.Println("Server starting at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
