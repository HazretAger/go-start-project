package main

import (
	"fmt"
	"go-start-project/db"
	"go-start-project/handler"
	"go-start-project/middleware"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database := db.Connect("localDB.sqlite")

    defer database.Close()

    db.InitSchema(database)

	// Проверяем корсы и авторизован ли пользователь
	http.HandleFunc("/user/register", middleware.CORS(handler.Register(database)))
	http.HandleFunc("/user/login", middleware.CORS(handler.Login(database)))
	http.HandleFunc("/user/getById", middleware.CORS(middleware.AuthCheck(handler.GetUserById(database))))
	http.HandleFunc("/user/getAllUsers", middleware.CORS(middleware.AuthCheck(handler.GetAllUsers(database))))

	fmt.Println("Server successfully started")
	http.ListenAndServe(":8080", nil)
}