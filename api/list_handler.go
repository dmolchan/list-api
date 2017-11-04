package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type List struct {
	ID          int
	Name        string
	Slug        string
	List        []byte
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ListHandler struct {
	DB *sql.DB
}

func (l ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := l.DB.Query(`
  SELECT id, name, slug, list, description, created_at, updated_at
  FROM lists
  `)
	if err != nil {
		fmt.Println(err)
		return
	}

	var lists []List

	for rows.Next() {
		l := List{}

		err = rows.Scan(&l.ID, &l.Name, &l.Slug, &l.List, &l.Description,
			&l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		lists = append(lists, l)
	}

	data, err := json.Marshal(lists)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}
