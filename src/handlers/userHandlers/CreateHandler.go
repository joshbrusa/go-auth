package userHandlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r)
		fmt.Println(r.Header)
		fmt.Println(r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusOK)
	}
}
