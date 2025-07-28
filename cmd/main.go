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

	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Protected())
	router.Use(middleware.WithDB(database))

	router.GET("/user/login", handler.Login)
	router.POST("/user/register", handler.Register)
	router.GET("/user/getById", handler.GetUserById)
	router.GET("/user/getAllUsers", handler.GetAllUsers)

	router.Run()
}