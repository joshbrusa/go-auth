package api

import (
	"github.com/gorilla/mux"

	v1 "github.com/joshbrusa/go-auth/src/api/v1/routers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	v1.UseRouter(router.PathPrefix("/v1").Subrouter())

	return router
}