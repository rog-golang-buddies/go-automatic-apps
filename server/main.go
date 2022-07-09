package server

import (
	"fmt"
	"net/http"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tables := strings.Join(getTables(), "<br/>")

		fmt.Fprintf(w, "Welcome to GAA! <br><br>"+tables)
	})

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Generate() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Header: "// GAA Generated",
		IDType: &field.TypeInfo{Type: field.TypeUUID},
	})
	if err != nil {
		panic(err)
	}
}

func getTables() []string {
	graph, err := entc.LoadGraph("./ent/schema", &gen.Config{
		Header: "// GAA Generated",
		IDType: &field.TypeInfo{Type: field.TypeUUID},
	})
	if err != nil {
		return []string{"Error", err.Error()}
	}

	result := []string{}
	tables, err := graph.Tables()
	if err != nil {
		return []string{"Error", err.Error()}
	}

	for _, table := range tables {
		result = append(result, table.Name)
	}
	return result
}
