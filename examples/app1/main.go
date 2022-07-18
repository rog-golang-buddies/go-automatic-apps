package main

import (
	gaa "github.com/rog-golang-buddies/go-automatic-apps/server"
)

func main() {
	var config gaa.ServerConfig = gaa.ServerConfig{
		Host: "localhost",
		Port: "8080",
	}
	gaa.Start(config)
}
