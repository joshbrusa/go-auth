package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joshbrusa/go-auth/src/routers"
	"github.com/joshbrusa/go-auth/src/utils"
)

func SetJsonSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@(mysql:3306)/db?parseTime=true")

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return db
}

func PingDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func StartServer(router *mux.Router) {
	port := utils.GetPort()

	slog.Info("Server listening on port: " + port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	// set up
	SetJsonSlog()

	// new db
	db := NewDB()
	PingDB(db)

	// new router
	router := routers.NewRouter(db)

	// start server
	StartServer(router)
}
