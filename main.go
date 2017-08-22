package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := NewRouter()

	// Adds cors middleware to router
	corsHandler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":3000", corsHandler))
}
