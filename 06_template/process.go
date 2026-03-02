package main

import (
	"html/template"
	"net/http"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "tmpl.html")
	t.ExecuteTemplate(w, "tmpl.html", "Hello World")
}
