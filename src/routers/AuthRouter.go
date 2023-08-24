package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/authHandlers"
)

func UseAuthRouter(authRouter *mux.Router) {
	authRouter.HandleFunc("/signIn", authHandlers.SignInHandler)
	authRouter.HandleFunc("/signOut", authHandlers.SignOutHandler)
	authRouter.HandleFunc("/signUp", authHandlers.SignUpHandler)
}