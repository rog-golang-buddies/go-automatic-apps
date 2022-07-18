package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

type ServerConfig struct {
	Host string
	Port string
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
		c.HTML(http.StatusOK, "config.tmpl", gin.H{
			"tables": database.GetTables(),
		})
	})

	router.Run(config.Host + ":" + config.Port)
}
