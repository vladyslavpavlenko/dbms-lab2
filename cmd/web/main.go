package main

import (
	"github.com/vladyslavpavlenko/dbms-lab2/internal/config"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	err := setup(&app)
	if err != nil {
		log.Fatal()
	}

	log.Printf("running on port %s...", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
