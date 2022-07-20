package server

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Host   string
	Port   string
	Tables []*schema.Table
}

//go:embed templates/*
var content embed.FS

func Start(config ServerConfig) {

	// Set defauls
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "8080"
	}

	fmt.Println("Starting server...")

	templates := template.Must(template.ParseFS(content, "templates/**"))
	router := gin.Default()
	router.SetHTMLTemplate(templates)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/config", func(c *gin.Context) {

		tables := []string{}
		for _, table := range config.Tables {
			tables = append(tables, table.Name)
		}

		c.HTML(http.StatusOK, "config.tmpl", gin.H{
			"tables": tables,
		})
	})

	err := router.Run(config.Host + ":" + config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		panic(err)
	}
}
