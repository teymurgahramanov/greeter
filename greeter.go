package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func GetAdjective() string {
	adjectives := []string{
		"Fantastic",
		"Gentle",
		"Suspisous",
		"Crazy",
		"Shy",
	}
	return adjectives[rand.Intn(len(adjectives))]
}

func GetName() string {
	names := []string{
		"Einstein",
		"Galilei",
		"Tesla",
		"Darwin",
		"Edison",
	}
	return names[rand.Intn(len(names))]
}

func GetIp(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func Greet(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Hi stranger from %s! My name is %s %s and I live on %s", GetIp(r), GetAdjective(), GetName(), hostname)
}

func main() {
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":8080", nil)
}
