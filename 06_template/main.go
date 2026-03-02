package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", Process)
	http.HandleFunc("/process2", Process2)
	http.HandleFunc("/process3", Process3)
	server.ListenAndServe()
}
