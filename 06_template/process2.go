package main

import (
	"html/template"
	"net/http"
)

func Process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "tmpl2.html")
	t.ExecuteTemplate(w, "tmpl2.html", "Hello World2")
}
