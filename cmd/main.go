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
	router.Use(middleware.WithDB(database))

	authorized := router.Group("/")

	authorized.Use(middleware.Protected())
	{
		authorized.GET("/user/getById", handler.GetUserById)
		authorized.GET("/user/getAllUsers", handler.GetAllUsers)
	}

	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)

	router.Run()
}