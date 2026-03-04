package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

func GobDemo() {
	type Post struct {
		Id      int
		Content string
		Author  string
	}
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "posts.gob")
	var postRead Post
	load(&postRead, "posts.gob")
	fmt.Println(postRead)
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
