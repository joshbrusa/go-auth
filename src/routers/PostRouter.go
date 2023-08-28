package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/postHandlers"
)

func UsePostRouter(router *mux.Router, db *sql.DB) {
	router.HandleFunc("", postHandlers.ReadAllHandler(db)).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", postHandlers.ReadHandler(db)).Methods("GET")
	router.HandleFunc("", postHandlers.CreateHandler(db)).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", postHandlers.UpdateHandler(db)).Methods("PATCH")
	router.HandleFunc("/{id:[0-9]+}", postHandlers.DeleteHandler(db)).Methods("DELETE")
}
