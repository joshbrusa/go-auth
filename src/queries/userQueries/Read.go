package userQueries

import (
	"database/sql"

	"github.com/joshbrusa/go-auth/src/models"
)

func SelectAll(db *sql.DB) []models.User {
	query := `
	SELECT * FROM user
	`
	rows, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdateAt,
			&user.Username,
			&user.Password,
		)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func SelectById(db *sql.DB, Id int) models.User {
	query := `
	SELECT * FROM user
	WHERE Id = ?
	`
	row := db.QueryRow(query, Id)

	var user models.User

	err := row.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdateAt,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		panic(err.Error())
	}

	return user
}