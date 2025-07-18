package model

import "time"

type User struct {
    ID int `json:"id"`  
    Name string `json:"name" validate:"required"`
    Surname string `json:"surname" validate:"required"`
    MiddleName string `json:"middle_name" validate:"required"`
    BirthDate time.Time `json:"birth_date" validate:"required"`
    PhoneNumber string `json:"phone_number" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
    IsVerified bool `json:"is_verified"`
    CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
    ID int `json:"id"`  
    Name string `json:"name" validate:"required"`
    Surname string `json:"surname" validate:"required"`
    MiddleName string `json:"middle_name" validate:"required"`
    BirthDate time.Time `json:"birth_date" validate:"required"`
    PhoneNumber string `json:"phone_number" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    IsVerified bool `json:"is_verified"`
}

type LoginResponse struct {
	Status int `json:"status"`
	Token string `json:"token"`
    User UserResponse `json:"user"`
}

type JWTPayload struct {
	Email string `json:"email"`
	Sub int `json:"sub"`
	IsVerified bool `json:"is_verified"`
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}