package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server started")
	http.ListenAndServe(":8000", nil)
}
