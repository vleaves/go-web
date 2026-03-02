package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func Process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}
