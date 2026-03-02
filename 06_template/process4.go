package main

import (
	"html/template"
	"net/http"
)

func Process4(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl4.html")
    daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    t.Execute(w, daysOfWeek)
}
