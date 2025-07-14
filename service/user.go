package service

import (
	"database/sql"
	"go-start-project/model"
)

func Register(db *sql.DB, user *model.User) error {
	_, err := db.Exec("INSERT INTO users (name, surname, middle_name, birth_date) VALUES (?, ?, ?, ?)", user.Name, user.Surname, user.MiddleName, user.BirthDate)

	if err != nil {
		return err
	}

    return nil
}

func GetAllUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]model.User, 0)

	for rows.Next() {
		var u model.User

		if err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.MiddleName, &u.BirthDate, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}