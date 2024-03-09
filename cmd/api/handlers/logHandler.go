package handlers

import (
	"encoding/json"
	"github.com/takshpanchal/log_ingestor/Internals/models"
	"io"
	"net/http"
)

func (app *Application) HandleIngest(w http.ResponseWriter, r *http.Request) {
	// only post requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO: Add auth

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.ErrLogger.Printf("Error: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log := models.Log{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		app.ErrLogger.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	app.ILogger.Printf("%+v \n", log)
}
