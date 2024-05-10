package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", portString)
	log.Fatal(srv.ListenAndServe())
}
