package userHandlers

import (
	"database/sql"
	"net/http"

	"github.com/joshbrusa/go-auth/src/queries/userQueries"
	"github.com/joshbrusa/go-auth/src/utils"
)

func ReadAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userQueries.SelectAll(db)
		
		response := utils.Response{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
			Data: users,
		}
		response.Send()
	}
}
