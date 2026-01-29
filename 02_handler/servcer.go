package main

import (
	"fmt"
	"net/http"
)

// type MyHandler struct{}

// func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

// func main() {
// 	mh := MyHandler{}

// 	server := http.Server{
// 		Addr:    ":8080",
// 		Handler: &mh,
// 	}

// 	server.ListenAndServe()
// }

type HelloHandler struct{}

func (hello *HelloHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type HiHandler struct{}

func (hi *HiHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hi World")
}

func main() {
	hello := HelloHandler{}
	hi := HiHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.Handle("/hello", &hello)
	http.Handle("/hi", &hi)
	server.ListenAndServe()
}
