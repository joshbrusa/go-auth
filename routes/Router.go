package routes

import (
	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	AuthRouter(*router.PathPrefix("/auth").Subrouter())
}