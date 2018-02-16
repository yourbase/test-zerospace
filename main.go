package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloWorldHandler struct{}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Root")
}

func main() {
	err := http.ListenAndServe(":8080", helloWorldHandler{})
	log.Fatal("HelloWorld ListenAndServe error", err)
}
