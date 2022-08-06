package config

import (
	"errors"

	"gorm.io/gorm"
)

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

func (c *ServerConfig) FindModel(modelName string) (ModelDescriptor, error) {
	for _, m := range c.Models {
		if m.Name == modelName {
			return m, nil
		}
	}
	return ModelDescriptor{}, errors.New("model [" + modelName + "] not found")
}
