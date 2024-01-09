package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
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
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies)
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error marshalling movies")
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	log.Info().Msg(tokens.Token)

	w.Write([]byte(tokens.Token))

}
