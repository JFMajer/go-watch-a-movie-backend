package main

import (
	"fmt"
	"net/http"
)

const port = 3000

type application struct {
	Domain string
}

func main() {
	app := &application {
		Domain: "example.com",
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		panic(err)
	}

}