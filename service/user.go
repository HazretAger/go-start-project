package service

import (
	"database/sql"
	"go-start-project/model"
)

func Register(db *sql.DB, user *model.User) error {
	_, err := db.Exec("INSERT INTO users (name, surname, middle_name, birth_date, phone_number, email, password, confirm_password, is_verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", user.Name, user.Surname, user.MiddleName, user.BirthDate, user.PhoneNumber, user.Email, user.Password, user.ConfirmPassword, user.IsVerified)

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

		if err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.MiddleName, &u.BirthDate, &u.PhoneNumber, &u.Email, &u.Password, &u.ConfirmPassword, &u.IsVerified, &u.CreatedAt); err != nil {
			return nil, err
		}
		
		users = append(users, u)
	}

	return users, nil
}

func GetUserHashedPassword(db *sql.DB, email string) (model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)

	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.MiddleName, &user.BirthDate, &user.PhoneNumber, &user.Email, &user.Password, &user.ConfirmPassword, &user.IsVerified, &user.CreatedAt); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func IsUserExists(db *sql.DB, email string) (bool, error) {
	var user model.User

	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)

	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.MiddleName, &user.BirthDate, &user.PhoneNumber, &user.Email, &user.Password, &user.ConfirmPassword, &user.IsVerified, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}