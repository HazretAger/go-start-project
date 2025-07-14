package model

import "time"

type User struct {
    ID        int        `json:"id"`
    Name     string     `json:"name"`
    Surname      string     `json:"surname"`
    MiddleName      string       `json:"middle_name"`
	BirthDate      time.Time       `json:"birth_date"`
    PhoneNumber    string       `json:"phone_number"`
    Email          string       `json:"email"`
    Password       string       `json:"password"`
    ConfirmPassword string       `json:"confirm_password"`
    IsVerified     bool         `json:"is_verified"`
    CreatedAt time.Time  `json:"created_at"`
}