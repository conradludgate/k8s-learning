package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://internal-api:8080/")
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		panic(err)
	}
	fmt.Fprintf(w, "internal api: %s\ntime: %s", body, time.Now())
}
