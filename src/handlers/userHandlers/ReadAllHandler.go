package userHandlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/joshbrusa/go-auth/src/models"
)

func ReadAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT ID, CreatedAt, UpdatedAt, Username, Password FROM user`
		rows, err := db.Query(query)

		if err != nil {
			panic(err.Error())
		}

		var users []models.User

		for rows.Next() {
			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.CreatedAt,
				&user.UpdateAt,
				&user.Username,
				&user.Password,
			)

			if err != nil {
				panic("Error scanning user rows.")
			}

			users = append(users, user)
		}

		json, err := json.Marshal(users)

		if err != nil {
			panic(err.Error())
		}

		w.Write(json)
	}
}
