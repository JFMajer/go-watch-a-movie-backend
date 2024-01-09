package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

const port = 3000

type application struct {
	Domain       string
	DB           repository.DatabaseRepo
	Logger       zerolog.Logger
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	log := configureLogger()
	dsn := "postgres://postgres:postgres@localhost:5432/movies?sslmode=disable&timezone=UTC"

	// Open a database connection pool
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Error().Err(err).Msg("Unable to create connection pool")
		os.Exit(1)
	}
	defer db.Close()

	// Check DB connection
	err = db.Ping(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Unable to ping database")
		os.Exit(1)
	}

	log.Info().Msg("Database connection successfully established")

	app := &application{
		Domain: "example.com",
		DB: &dbrepo.PostgresDBRepo{
			DB:     db,
			Logger: log,
		},
		Logger:       log,
		JWTSecret:    "verysecret",
		JWTIssuer:    "example.com",
		JWTAudience:  "example.com",
		CookieDomain: "localhost",
	}

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	log.Info().Msgf("Server is starting on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		panic(err)
	}

}
