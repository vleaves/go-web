package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	Db, err := sql.Open("sqlite", "posts.db")
	if err != nil {
		panic(err)
	}
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		author TEXT NOT NULL
	);`
	if _, err := Db.Exec(createTableSQL); err != nil {
		log.Fatal("建表失败", err)
	} else {
		fmt.Println("建表成功")
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit ?", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values (?, ?)"
	_, err = Db.Exec(statement, post.Content, post.Author)
	if err != nil {
		return fmt.Errorf("插入失败: %w", err)
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return fmt.Errorf("获取ID失败: %w", err)
	// }
	// post.Id = int(id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id= ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau sheong"}
	fmt.Println(post)
	post.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	// readPost.Content = "Bonjour Monde!"
	// readPost.Author = "Pierre"
	// if err := readPost.Update(); err != nil {
	// 	log.Fatal("更新失败:", err)
	// }
	// fmt.Println("更新后:", readPost)
	// posts, err := Posts(10) // 显式传入 limit
	// if err != nil {
	// 	log.Fatal("查询列表失败:", err)
	// }
	// fmt.Println("当前帖子列表:", posts)

	// // 删除
	// if err := readPost.Delete(); err != nil {
	// 	log.Fatal("删除失败:", err)
	// }
	// fmt.Println("帖子已删除")
}
