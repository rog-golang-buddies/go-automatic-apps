package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql/schema"
)

type ServerConfig struct {
	Host   string
	Port   string
	Tables []*schema.Table
}

//go:embed web/dist
var webDistEmbed embed.FS

func Start(config ServerConfig) {

	// Set defauls
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "8080"
	}

	fmt.Println("Starting server...")

	webRoot, err := fs.Sub(webDistEmbed, "web/dist")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(webRoot)))

	var serverUrl = config.Host + ":" + config.Port
	log.Println("Server started at " + serverUrl)
	err = http.ListenAndServe(serverUrl, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
