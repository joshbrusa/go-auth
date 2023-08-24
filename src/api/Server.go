package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joshbrusa/go-auth/src/utils"
)


func StartServer() {
	port := utils.GetPort()
	router := NewRouter()

	slog.Info("Server listening on port " + port)

	err := http.ListenAndServe(port, router)

	if (err != nil) {
		slog.Error(err.Error())
		os.Exit(1)
	}
}