package main

import (
	"fmt"
	"net/http"
)

func main() {
	StartServer()
}

func StartServer() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	data := []byte("Hello, World!")
	w.Write(data)
}
