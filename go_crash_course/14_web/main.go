package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Fprint(w, "<h1>Hello From Server</h1>")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello From Server</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	fmt.Println("Server connected")
	http.ListenAndServe(":3000", nil)
}
