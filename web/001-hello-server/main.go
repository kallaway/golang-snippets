package main

// Run this file, then go to localhost:8080.
// Type in localhost:8080/YourName
// You will see the server print out "Hello YourName" on the page

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Path
	m = strings.TrimPrefix(m, "/")
	m = "Hello " + m

	// convert to a slice of bytes for the Writer
	w.Write([]byte(m))
}

func main() {
	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
