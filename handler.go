package main

import (
	"io"
	"net/http"
)

func (app *application) handleIngest(w http.ResponseWriter, r *http.Request) {
	// only post requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO: Add auth

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.logger.Printf("Error: %s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	app.logger.Println(string(body))
}
