package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/rootHandlers"
)

func UseRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandlers.RootHandler)

	authRouter := router.PathPrefix("/auth").Subrouter()
	UseAuthRouter(authRouter)

	return router
}