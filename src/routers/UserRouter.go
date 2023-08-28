package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/userHandlers"
)

func UseUserRouter(router *mux.Router, db *sql.DB) {
	router.HandleFunc("", userHandlers.ReadAllHandler(db)).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", userHandlers.ReadHandler(db)).Methods("GET")
	router.HandleFunc("", userHandlers.CreateHandler(db)).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", userHandlers.UpdateHandler(db)).Methods("PATCH")
	router.HandleFunc("/{id:[0-9]+}", userHandlers.DeleteHandler(db)).Methods("DELETE")
}
