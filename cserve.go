package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Hello world!")
	io.WriteString(w, "<pre>")
	io.WriteString(w, "Hello world!")
}

func main() {
	//http.HandleFunc("/", hello)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":80", nil)
}
