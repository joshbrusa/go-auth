package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/api/v1/handlers/rootHandlers"
)

func UseRouter(router *mux.Router) {
	router.HandleFunc("/", rootHandlers.RootHandler)

	UseAuthRouter(router.PathPrefix("/auth").Subrouter())
}