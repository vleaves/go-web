package main

import (
	"fmt"
	"net/http"
)

func Form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}
