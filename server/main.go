package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tables := strings.Join(database.GetTables(), "<br/>")

		fmt.Fprintf(w, "<html><body>Welcome to GAA! <br><br>"+tables+"</body></html>")
	})

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
