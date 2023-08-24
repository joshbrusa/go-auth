package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/rootHandler"
)

func UseRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler.Root)

	authRouter := router.PathPrefix("/auth").Subrouter()
	UseAuthRouter(authRouter)

	return router
}