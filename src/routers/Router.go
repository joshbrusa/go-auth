package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/rootHandlers"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandlers.RootHandler())

	authRouter := router.PathPrefix("/auth").Subrouter()
	UseAuthRouter(authRouter, db)

	userRouter := router.PathPrefix("/user").Subrouter()
	UseUserRouter(userRouter, db)

	postRouter := router.PathPrefix("/post").Subrouter()
	UseUserRouter(postRouter, db)

	return router
}
