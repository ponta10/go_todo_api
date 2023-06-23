package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースへの接続情報
    db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// INSERT文を準備
	stmt, err := db.Prepare("INSERT INTO todos (task) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// データを追加するタスクのリスト
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	// タスクをループしてデータを追加
	for _, task := range tasks {
		_, err := stmt.Exec(task)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Seeder executed successfully")
}
