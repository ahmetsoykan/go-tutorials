package main

import (
	"fmt"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ok!")
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":80", nil)
}
