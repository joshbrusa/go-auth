package userHandlers

import (
	"database/sql"
	"net/http"

	"github.com/joshbrusa/go-auth/src/utils"
)

func ReadHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expectedItems []string

		expectedItems = append(expectedItems, "Id")
		
		ok := utils.CheckRequestBody(r, expectedItems)

		if !ok {
			// if not ok
		}

		w.WriteHeader(http.StatusOK)

		// user := userQueries.SelectById(db, 1)

		// utils.WriteJson(w, user)
	}
}
