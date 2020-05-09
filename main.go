package main

import (
	"fmt"
	"net/http"
)

func main() {
	// the handler func handles the page to display
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Go HTTP")
	})

	fmt.Println("Server listening on 8000")
	http.ListenAndServe(":8000", http.DefaultServeMux)
}
