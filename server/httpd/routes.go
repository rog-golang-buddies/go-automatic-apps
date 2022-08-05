package httpd

import (
	"log"
	"net/http"

	"github.com/rog-golang-buddies/go-automatic-apps/config"
)

func CreateGetModelsHandler(config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {

	// Get model names
	modelNames := make([]string, 0)
	for _, model := range config.Models {
		modelNames = append(modelNames, model.Name)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := WriteJSON(w, http.StatusOK, modelNames)
		if err != nil {
			log.Fatalf("Error on GetModels: %s", err.Error())
		}
	}
}
