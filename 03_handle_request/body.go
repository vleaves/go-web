package main

import (
	"fmt"
	"net/http"
)

func Body(writer http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(writer, string(body))
}
