package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AMFDPMTE/list"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Hello World")
	l := list.New()
	fmt.Println(l.Serialize())

	// sqlite time!
	db, err := sql.Open("sqlite3", "./db/db.sqlite3")
	if err != nil {
		panic(err)
	}

	// insert some data
	stmt, err := db.Prepare(`
	INSERT INTO lists(name, slug, list, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		name, slug, listBlob, createdAt, updatedAt := uuid.NewV4().String(),
			uuid.NewV4().String(), []byte{}, time.Now().UTC(),
			time.Now().UTC()
		result, err := stmt.Exec(name, slug, listBlob, createdAt, updatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

	// select
	stmt, err = db.Prepare("SELECT * FROM lists")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		fmt.Println(rows.Columns())
	}
}
