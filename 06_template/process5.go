package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func Process5(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl5.html").Funcs(funcMap)
	t.ParseFiles("tmpl5.html")
	t.Execute(w, time.Now())
}
