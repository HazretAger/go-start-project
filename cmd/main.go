package main

import (
	"fmt"
	"go-start-project/db"
	"go-start-project/handler"
	"net/http"
)

func main() {
	database := db.Connect("localDB.sqlite")

    defer database.Close()

    db.InitSchema(database)

	// Auth routes
	http.HandleFunc("/register", db.CORS(handler.Register(database)))
	http.HandleFunc("/getUsers", db.CORS(handler.GetAllUsers(database)))

	fmt.Println("Сервер запущен")
	http.ListenAndServe(":8080", nil)
}