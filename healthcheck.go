package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"enviornment": app.cfg.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.log.Println(err)
		http.Error(w, "The serve encountered unexpected problrm", http.StatusInternalServerError)
	}
}
