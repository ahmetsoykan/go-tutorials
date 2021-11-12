package main

import (
	"log"
	"net/http"

	"github.com/ahmetsoykan/go-tutorials/047-factoriel/cmd/internal/handlers"
)

func main() {

	log.Printf("main : Started")

	api := http.Server{
		Addr:    "localhost:8000",
		Handler: handlers.API(),
	}

	if err := api.ListenAndServe(); err != nil {
		log.Fatal("error: listeing and serving: %s", err)
	}
}
