package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	docker_namegenerator "github.com/docker/docker/pkg/namesgenerator"
)

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
		os.Exit(1)
	}
	fmt.Fprintf(w, "Hi stranger from %s! My name is %s and I live on %s", GetIp(r), strings.ToTitle(docker_namegenerator.GetRandomName(0)), hostname)
}

func main() {
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":8080", nil)
}
