package main

import (
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var id string

func main() {
	id = uuid.New().String()

	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, id)
}
