package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

const port = 3000

type application struct {
	Domain string
	DB     *pgxpool.Pool
	Logger zerolog.Logger
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
		panic(err)
	}

	log.Info().Msg("Database connection successfully established")

	app := &application{
		Domain: "example.com",
		DB:     db,
		Logger: log,
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		panic(err)
	}

}
