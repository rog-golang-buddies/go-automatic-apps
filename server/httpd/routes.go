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

func CreateGetModelRows(config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		modelName, ok := ctx.Value("model").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		var model any
		for _, m := range config.Models {
			if m.Name == modelName {
				model = m.Model
				break
			}
		}

		config.DB.Find(model)

		err := WriteJSON(w, http.StatusOK, "yeah")
		if err != nil {
			log.Fatalf("Error on GetModelRows: %s", err.Error())
		}
	}
}
