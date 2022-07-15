package database

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func GetTables() []string {
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
