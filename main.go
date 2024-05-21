package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/adamhu714/rssagg/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Error getting PORT environment variable")
	}

	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("Error getting DB_URL environment variable")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUserCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUserGet))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))

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
