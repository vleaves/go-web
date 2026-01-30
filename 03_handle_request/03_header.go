package main

import (
	"fmt"
	"net/http"
)

func Headers(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request.Header)
}
