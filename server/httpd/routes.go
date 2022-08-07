package httpd

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rog-golang-buddies/go-automatic-apps/config"
	"gorm.io/gorm/schema"
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

	type FieldInfo struct {
		Name string
		Type string
		Size int
	}

	type RowsResult struct {
		ModelName string
		TableName string
		Fields    []FieldInfo
		Data      []map[string]interface{}
	}

	cache := &sync.Map{}
	namer := config.DB.NamingStrategy

	return func(w http.ResponseWriter, r *http.Request) {

		modelName := chi.URLParam(r, "model")

		model, err1 := config.FindModel(modelName)
		if err1 != nil {
			log.Fatalf("Error on GetModelRows: %s", err1.Error())
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(err1.Error()))
			if err != nil {
				log.Fatalf("Error on GetModelRows: %s", err.Error())
			}
			return
		}

		result := RowsResult{}

		schema, errSchema := schema.Parse(model.Model, cache, namer)
		if errSchema != nil {
			log.Fatalln("error with schema", errSchema.Error())
		} else {
			log.Println("schema", schema)
			result.ModelName = schema.ModelType.Name()
			result.TableName = schema.Table

			fields := []FieldInfo{}
			for _, f := range schema.Fields {
				fi := FieldInfo{
					Name: f.Name,
					Type: f.FieldType.Name(),
					Size: f.Size,
				}
				fields = append(fields, fi)
			}
			result.Fields = fields
		}

		data := []map[string]interface{}{}
		config.DB.
			Model(model.Model).
			Find(&data)

		result.Data = data

		err := WriteJSON(w, http.StatusOK, result)
		if err != nil {
			log.Fatalf("Error on GetModelRows: %s", err.Error())
		}
	}
}
