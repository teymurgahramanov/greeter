package main

import (
	"fmt"
	"net/http"
	"os"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Hi! I am on %s", hostname)
}

func main() {
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":8080", nil)
}
