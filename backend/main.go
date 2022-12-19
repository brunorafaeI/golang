package main

import (
	"fmt"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!!")
	// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {

	http.HandleFunc("/", MainHandler)

	fmt.Println("Listening on port 5050...")
	http.ListenAndServe(":5050", nil)
}
