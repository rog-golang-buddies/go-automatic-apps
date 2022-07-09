package server

import (
	"fmt"
	"net/http"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
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

func Generate() {
	entc.Generate("./ent/schema", &gen.Config{
		Header: "// GAA Generated",
		IDType: &field.TypeInfo{Type: field.TypeUUID},
	})
}
