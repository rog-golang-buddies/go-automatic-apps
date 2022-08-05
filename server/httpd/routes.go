package httpd

import (
	"net/http"

	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

func (c *controller) GetModels(w http.ResponseWriter, r *http.Request) {
	tables := database.GetTables()
	err := WriteJSON(w, http.StatusOK, tables)
	if err != nil {
		log.Fatalf("Error on GetModels: %s", err.Error())
	}
}
