package main

import (
	"io"
	"net/http"

	"github.com/google/uuid"
)

var id string

func main() {
	id = uuid.New().String()

	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, id)
}
