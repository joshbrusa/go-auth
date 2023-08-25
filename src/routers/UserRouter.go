package routers

import (
	"github.com/gorilla/mux"
	userhandlers "github.com/joshbrusa/go-auth/src/handlers/userHandlers"
)

func UseUserRouter(router *mux.Router) {
	router.HandleFunc("/create", userhandlers.CreateHandler)
	router.HandleFunc("/read", userhandlers.ReadHandler)
	router.HandleFunc("/update", userhandlers.UpdateHandler)
	router.HandleFunc("/delete", userhandlers.DeleteHandler)
}
