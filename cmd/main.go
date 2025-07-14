package main

import (
	"fmt"
	"go-start-project/db"
	"go-start-project/handler"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database := db.Connect("localDB.sqlite")

    defer database.Close()

    db.InitSchema(database)

	// Auth routes
	http.HandleFunc("/register", db.CORS(handler.Register(database)))
	http.HandleFunc("/login", db.CORS(handler.Login(database)))
	http.HandleFunc("/getUsers", db.CORS(handler.GetAllUsers(database)))

	fmt.Println("Server successfully started")
	http.ListenAndServe(":8080", nil)
}