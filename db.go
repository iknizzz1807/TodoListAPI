package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./todos.db")
    if err != nil {
        panic(err)
    }

    sqlStatement := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        finished BOOLEAN
    );
    `
    _, err = db.Exec(sqlStatement)
    if err != nil {
        panic(err)
    }
}