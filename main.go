package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

type Todo struct {
    ID    int    `json:"id"`
    Task  string `json:"task"`
}

var db *sql.DB
var err error

func main() {
    db, err = sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := mux.NewRouter()

    router.HandleFunc("/todos", getTodos).Methods("GET")
    router.HandleFunc("/todos", createTodo).Methods("POST")
    
    fmt.Println("Server starting at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var todos []Todo

    result, err := db.Query("SELECT id, task FROM todos")
    if err != nil {
        panic(err.Error())
    }

    defer result.Close()

    for result.Next() {
        var todo Todo
        err := result.Scan(&todo.ID, &todo.Task)
        if err != nil {
            panic(err.Error())
        }
        todos = append(todos, todo)
    }

    json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var todo Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    stmt, err := db.Prepare("INSERT INTO todos (task) VALUES (?)")
    if err != nil {
        panic(err.Error())
    }

    res, err := stmt.Exec(todo.Task)
    if err != nil {
        panic(err.Error())
    }

    id, err := res.LastInsertId()
    if err != nil {
        panic(err.Error())
    }

    todo.ID = int(id)

    fmt.Println(todo)

    json.NewEncoder(w).Encode(todo)
}



