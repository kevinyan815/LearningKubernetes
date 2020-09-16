package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s", hostname)
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}