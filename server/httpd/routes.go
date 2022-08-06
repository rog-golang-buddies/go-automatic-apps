package httpd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func CreateGetModelRows(config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		modelName := chi.URLParam(r, "model")

		model, err1 := config.FindModel(modelName)
		if err1 != nil {
			log.Fatalf("Error on GetModelRows: %s", err1.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err1.Error()))
			return
		}

		result := []map[string]interface{}{}
		config.DB.
			Model(model.Model).
			Find(&result)

		err := WriteJSON(w, http.StatusOK, result)
		if err != nil {
			log.Fatalf("Error on GetModelRows: %s", err.Error())
		}
	}
}
