package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome GAA!")
	})

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
