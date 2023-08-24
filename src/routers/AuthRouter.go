package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/authHandler"
)

func UseAuthRouter(authRouter *mux.Router) {
	authRouter.HandleFunc("/signIn", authHandler.SignIn)
	authRouter.HandleFunc("/signUp", authHandler.SignUp)
	authRouter.HandleFunc("/signOut", authHandler.SignOut)
}