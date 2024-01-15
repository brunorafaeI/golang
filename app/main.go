package main

import (
	"fmt"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello, world!!")
	if err != nil {
		return
	}
	// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {

	http.HandleFunc("/", MainHandler)
	fmt.Println("Listening on port 5050...")
	_, err := http.ListenAndServe(":5050", nil)

	if err != nil {
		return
	}
}
