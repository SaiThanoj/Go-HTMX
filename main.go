package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, HTMX with Go!")
		http.ListenAndServe(":3000", nil)
	})
}
