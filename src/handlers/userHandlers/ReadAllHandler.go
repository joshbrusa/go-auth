package userhandlers

import (
	"database/sql"
	"net/http"
)

func ReadAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
