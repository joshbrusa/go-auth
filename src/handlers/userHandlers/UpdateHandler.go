package userHandlers

import (
	"database/sql"
	"net/http"
)

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
