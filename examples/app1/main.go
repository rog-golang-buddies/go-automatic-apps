package main

import (
	"app1/models"

	cfg "github.com/rog-golang-buddies/go-automatic-apps/config"
	gaa "github.com/rog-golang-buddies/go-automatic-apps/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var config = cfg.ServerConfig{
		Host:     "localhost",
		HttpPort: "8080",
		DB:       db,
		Models: []cfg.ModelDescriptor{
			cfg.ModelDescriptor{
				Name:  "Product",
				Model: models.Product{},
			},
		},
	}

	gaa.Start(config)
}
