package main

import (
	"go-start-project/db"
	"go-start-project/handler"
	"go-start-project/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database := db.Connect("localDB.sqlite")

    defer database.Close()

    db.InitSchema(database)

	// Проверяем корсы и авторизован ли пользователь
	// http.HandleFunc("/user/register", middleware.CORS(handler.Register(database)))
	// http.HandleFunc("/user/getById", middleware.CORS(middleware.Protected(handler.GetUserById(database))))
	// http.HandleFunc("/user/getAllUsers", middleware.CORS(middleware.Protected(handler.GetAllUsers(database))))

	// fmt.Println("Server successfully started")
	// http.ListenAndServe(":8080", nil)


	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.WithDB(database))

	router.GET("/user/login", handler.Login)

	router.Run()
}