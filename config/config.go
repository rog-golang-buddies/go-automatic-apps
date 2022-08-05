package config

import "gorm.io/gorm"

type ModelDescriptor struct {
	Name  string
	Model interface{}
}

type ServerConfig struct {
	Host     string
	HttpPort string
	DB       *gorm.DB
	Models   []ModelDescriptor
}
