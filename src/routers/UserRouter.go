package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	userhandlers "github.com/joshbrusa/go-auth/src/handlers/userHandlers"
)

func UseUserRouter(router *mux.Router, db *sql.DB) {
	router.HandleFunc("", userhandlers.ReadAllHandler(db)).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", userhandlers.ReadHandler(db)).Methods("GET")
	router.HandleFunc("", userhandlers.CreateHandler(db)).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", userhandlers.UpdateHandler(db)).Methods("PATCH")
	router.HandleFunc("/{id:[0-9]+}", userhandlers.DeleteHandler(db)).Methods("DELETE")
}
