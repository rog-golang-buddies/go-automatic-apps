package httpd

import (
	"net/http"

	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

func (c *controller) GetModels(w http.ResponseWriter, r *http.Request) {
	tables := database.GetTables()
	WriteJSON(w, http.StatusOK, tables)
}
