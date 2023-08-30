package userHandlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/joshbrusa/go-auth/src/types"
)

func ReadAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// testing
		panic("this is a test")
		
		query := `SELECT id, username, password, created_at FROM users`
		rows, err := db.Query(query)

		if (err != nil) {
			w.WriteHeader(http.StatusInternalServerError)
		}

		var users []types.User

		for rows.Next() {
			var user types.User

			err := rows.Scan(
				&user.CreatedAt,
				&user.UpdateAt,
				&user.ID,
				&user.Username,
				&user.Password,
			)

			if (err != nil) {
				w.WriteHeader(http.StatusInternalServerError)
			}

			users = append(users, user)
		}

		json, err := json.Marshal(users)

		if (err != nil) {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(json)
	}
}
