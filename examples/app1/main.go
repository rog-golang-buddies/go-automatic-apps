package main

import (
	"app1/ent"
	"app1/ent/migrate"
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	gaa "github.com/rog-golang-buddies/go-automatic-apps/server"
)

func main() {

	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var tables = migrate.Tables

	var config gaa.ServerConfig = gaa.ServerConfig{
		Host:   "localhost",
		Port:   "8080",
		Tables: tables,
	}

	gaa.Start(config)
}
