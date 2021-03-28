package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world.")
}

func main() {
	fmt.Println("Starting Server at http://localhost:8080")
	http.HandleFunc("/user/get", handler)
	http.ListenAndServe(":8080", nil)
}
