package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":8080", nil)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
