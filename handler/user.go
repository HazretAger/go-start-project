package handler

import (
	"database/sql"
	"encoding/json"
	"go-start-project/model"
	"go-start-project/service"
	"go-start-project/utils"
	"net/http"

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

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка метода запроса
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var logoPass model.User

		if err := json.NewDecoder(r.Body).Decode(&logoPass); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Получение пользователя по email
		user, err := service.GetUserHashedPassword(db, logoPass.Email)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Сравнение пароля
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logoPass.Password))

		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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