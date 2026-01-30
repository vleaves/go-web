package main

import "net/http"

func main() {
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/body", Body)
	http.HandleFunc("/form", Form)
	http.ListenAndServe(":8080", nil)
}
