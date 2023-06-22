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

// w: レスポンス, r: リクエストヘッダ
func createTodo(w http.ResponseWriter, r *http.Request) {
    // Json形式で返す
    w.Header().Set("Content-Type", "application/json")

    // Todo構造体でtodoを定義
    // ここでは{0, 空文字}の初期値が入っている
    var todo Todo

    // リクエストボディのJSONをtodoの構造体の型式に変換
    // ここでは{0, タスク}と送られてきたtaskが入る
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // todoに格納されたTaskをtososテーブルに挿入
    res, err := db.Exec("INSERT INTO todos (task) VALUES (?)", todo.Task)
    if err != nil {
        panic(err.Error())
    }
    
    // 前のINSERT操作で新たに追加された行のIDを取得し、idに入れる
    id, err := res.LastInsertId()
    if err != nil {
        panic(err.Error())
    }

    // todoのIDに先ほどのidを入れる
    todo.ID = int(id)

    // Jsonコードに変換して返却
    json.NewEncoder(w).Encode(todo)
}



