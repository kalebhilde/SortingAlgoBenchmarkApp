package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, requester *http.Request) {
	fmt.Fprintln(writer, "Backend is running inside Docker!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on: 8080")
	http.ListenAndServe(":8080", nil)
}