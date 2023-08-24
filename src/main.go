package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshbrusa/go-auth/src/api"
)

func SetJsonSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
	slog.SetDefault(logger)
}

func LoadEnv() {
	err := godotenv.Load()

	if (err != nil) {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	SetJsonSlog()
	LoadEnv()
	api.StartServer()
}
