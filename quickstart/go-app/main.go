package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, you are visiting host: %s", hostname)
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server starting on port:8080...")
	http.ListenAndServe(":8080", nil)
}
