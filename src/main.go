package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"go-auth/src/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	loadErr := godotenv.Load()

  if loadErr != nil {
    log.Fatal("Error loading .env file.")
  }

  port := os.Getenv("PORT")
	
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/auth", handlers.AuthHandler)

	slog.Info("Server listening on port " + port)

	ListenAndServeErr := http.ListenAndServe(port, r)

	if ListenAndServeErr != nil {
		log.Fatal("Error listening and serving.")
	}
}
