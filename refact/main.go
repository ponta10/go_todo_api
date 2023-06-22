package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/ponta10/go_todo_api/refact/handlers"
    "github.com/ponta10/go_todo_api/refact/models"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    todoModel := models.NewTodoModel(db)
    todoHandler := handlers.NewTodoHandler(todoModel)

    router := mux.NewRouter()
    router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
    router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

    fmt.Println("Server starting at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
