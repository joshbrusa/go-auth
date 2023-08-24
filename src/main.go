package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshbrusa/go-auth/src/routers"
	"github.com/joshbrusa/go-auth/src/utils"
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

func StartServer() {
	port := utils.GetPort()
	router := routers.UseRouter()

	slog.Info("Server listening on port " + port)

	err := http.ListenAndServe(port, router)

	if (err != nil) {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	SetJsonSlog()
	LoadEnv()
	StartServer()
}
