package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

//go:embed templates/*
var content embed.FS

func Start() {

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

	fmt.Println("Server started on http://localhost:8080")
	router.Run()
}
