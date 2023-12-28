package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "go movies up and running",
		Version: "1.0.0",
	}

	err := app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error writing JSON")
		return
	}

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error retrieving all movies")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies)
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error marshalling movies")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
