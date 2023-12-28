package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

const port = 3000

type application struct {
	Domain string
	DB     *pgxpool.Pool
}

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/movies?sslmode=disable&timezone=UTC"

	// Open a database connection
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check DB connection
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	app := &application{
		Domain: "example.com",
		DB:     db,
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		panic(err)
	}

}
