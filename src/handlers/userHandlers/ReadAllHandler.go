package userHandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/joshbrusa/go-auth/src/models"
	"github.com/joshbrusa/go-auth/src/utils"
)

func ReadAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, file, line, _ := runtime.Caller(0)

		panicData := utils.PanicData{
			Msg: "Error encoding json.",
			File: file,
			Line: line,
		}
		utils.NewPanic(panicData)
		panic("Error encoding json.")

		query := `SELECT id, username, password, created_at FROM users`
		rows, err := db.Query(query)

		if (err != nil) {
			panic("Error querying users.")
		}

		var users []models.User

		for rows.Next() {
			var user models.User

			err := rows.Scan(
				&user.CreatedAt,
				&user.UpdateAt,
				&user.ID,
				&user.Username,
				&user.Password,
			)

			if (err != nil) {
				panic("Error scanning user rows.")
			}

			users = append(users, user)
		}

		json, err := json.Marshal(users)

		if (err != nil) {
			pc, file, line, ok := runtime.Caller(1000)
			fmt.Println(pc)
			fmt.Println(file)
			fmt.Println(line)
			fmt.Println(ok)
			panic("Error encoding json.")
		}

		w.Write(json)
	}
}
