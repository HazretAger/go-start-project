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

func GetAllUsers(db *sql.DB) ([]model.UserResponse, error) {
	rows, err := db.Query("SELECT id, name, surname, middle_name, birth_date, phone_number, email, is_verified FROM users")
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]model.UserResponse, 0)

	for rows.Next() {
		var u model.UserResponse

		if err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.MiddleName, &u.BirthDate, &u.PhoneNumber, &u.Email, &u.IsVerified); err != nil {
			return nil, err
		}
		
		users = append(users, u)
	}

	return users, nil
}

func GetUserByEmail(db *sql.DB, email string) (model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)

	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.MiddleName, &user.BirthDate, &user.PhoneNumber, &user.Email, &user.Password, &user.ConfirmPassword, &user.IsVerified, &user.CreatedAt); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GetUserById(db *sql.DB, id int) (model.UserResponse, error) {
	var user model.UserResponse

	row := db.QueryRow("SELECT id, name, surname, middle_name, birth_date, phone_number, email, is_verified FROM users WHERE id=?", id)

	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.MiddleName, &user.BirthDate, &user.PhoneNumber, &user.Email, &user.IsVerified); err != nil {
		return model.UserResponse{}, err
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