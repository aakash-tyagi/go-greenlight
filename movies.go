package main

import (
	"fmt"
	"net/http"
	"time"

	data "github.com/aakash-tyagi/greenlight/internal"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Date(2012, time.January, 31, 22, 22, 22, 22, time.Now().Location()),
		Title:     "Gundey",
		Runtime:   3,
		Genres:    []string{"comedy"},
		Version:   1,
	}

	err = app.writeJSON(w, 200, movie, http.Header{})
	if err != nil {
		app.log.Println(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
