package handler

import (
	"database/sql"
	"go-start-project/model"
	"go-start-project/service"
	"go-start-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var user model.User

	// Декодирование данных пользователя из тела запроса
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Decoding JSON error	",
		})
		return
	}

	isUserExists, _ := service.IsUserExists(db, user.Email)

	if isUserExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "User already exists",
		})
		return
	}

	// Валидация данных пользователя
	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Validate error",
		})
		return
	}

	// Хеширование пароля
	hashedPass, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Error with password hashing",
		})
		return
	}

	user.Password = hashedPass
	
	// Регистрация пользователя
	if err := service.Register(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Register error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"message": "User created",
	})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	var logoPass model.Login

	if err := c.ShouldBindBodyWithJSON(&logoPass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Incorrect JSON",
		})
		return
	}

	// Получение пользователя по email
	user, err := service.GetUserByEmail(db, logoPass.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	// Генерация access и refresh токена
	tokens, err := utils.GetAccessAndRefreshTokens(model.JWTPayload{
		Sub: user.ID,
		Email: user.Email,
		IsVerified: user.IsVerified,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Error with token",
		})
		return
	}

	// Сравнение пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logoPass.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": "User's unauthorized",
		})
		return
	}

	// Установка refresh токена в cookie
	c.SetCookie(
		"refresh_token", 
		tokens.RefreshToken,
		30 * 24 * 60 * 60, 
		"/",
		"",
		false,
		true,
	)

	// Отправка данных пользователя клиенту
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": http.StatusOK,
		"token": tokens.AccessToken,
		"user": model.UserResponse{
			ID: user.ID,
			Email: user.Email,
			Name: user.Name,
			Surname: user.Surname,
			MiddleName: user.MiddleName,
			BirthDate: user.BirthDate,
			PhoneNumber: user.PhoneNumber,
			IsVerified: user.IsVerified,
		},
	})
}

func GetUserById(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	db := c.MustGet("db").(*sql.DB)

	// Получаем ID и сразу же преобразуем его в формат int
	id, _ := strconv.Atoi(c.Query("id"))

	user, err := service.GetUserById(db, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Decoding JSON error	",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	// Получение всех пользователей
	users, err := service.GetAllUsers(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	// Отправка данных пользователей клиенту
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"users": users,
	})
}