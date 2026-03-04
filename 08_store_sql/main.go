package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // 替换为 SQLite 驱动
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	Db, err := sql.Open("sqlite3", "file:posts.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		author TEXT NOT NULL
	);`
	if _, err : Db.Exec(createTableSQL); err != nil {
		log.Fatal("建表失败", err)
	}
}

func Posts(limit int) (posts []Post, err error){
	rows, err := Db.Query("select id, content, author from posts limit ?", limit)
	if err != nil {
		return
	}
	for rows.Next(){
		post := Post{}
		err := rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
