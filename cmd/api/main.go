package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const port = 3000

type application struct {
	Domain string
	DB     *sql.DB
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC"

	// Open a database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check DB connection
	err = db.Ping()
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
