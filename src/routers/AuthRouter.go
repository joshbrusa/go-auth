package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/authHandlers"
)

func UseAuthRouter(router *mux.Router, db *sql.DB) {
	router.Handle("/signIn", authHandlers.SignInHandler(db))
	router.HandleFunc("/signOut", authHandlers.SignOutHandler(db))
	router.HandleFunc("/signUp", authHandlers.SignUpHandler(db))
}
