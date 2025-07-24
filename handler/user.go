package handler

import (
	"database/sql"
	"encoding/json"
	"go-start-project/model"
	"go-start-project/service"
	"go-start-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка метода запроса
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var user model.User

		// Декодирование данных пользователя из тела запроса
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Проверка что пользователь с таким email уже существует
		isUserExists, _ := service.IsUserExists(db, user.Email)

		if isUserExists {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Валидация данных пользователя
		validate := validator.New()
		err := validate.Struct(user)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Хеширование пароля
		hashedPass, err := utils.HashPassword(user.Password)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		user.Password = hashedPass
		
		// Регистрация пользователя
		if err := service.Register(db, &user); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
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

func GetUserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Проверка метода запроса
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		// Получаем ID и сразу же преобразуем его в формат int
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		user, err := service.GetUserById(db, id)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

		json.NewEncoder(w).Encode(user)
	}
}

func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Проверка метода запроса
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		// Получение всех пользователей
		users, err := service.GetAllUsers(db)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Отправка данных пользователей клиенту
		json.NewEncoder(w).Encode(users)
	}
}