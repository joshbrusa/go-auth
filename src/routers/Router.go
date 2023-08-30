package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/handlers/rootHandlers"
	"github.com/joshbrusa/go-auth/src/middleware"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	router.Use(middleware.PanicMiddleware)

	router.HandleFunc("/", rootHandlers.RootHandler())

	authRouter := router.PathPrefix("/auth").Subrouter()
	UseAuthRouter(authRouter, db)

	userRouter := router.PathPrefix("/user").Subrouter()
	UseUserRouter(userRouter, db)

	postRouter := router.PathPrefix("/post").Subrouter()
	UseUserRouter(postRouter, db)

	return router
}
