package main

import "net/http"

func main() {
	http.HandleFunc("/headers", Headers)
	http.ListenAndServe(":8080", nil)
}
