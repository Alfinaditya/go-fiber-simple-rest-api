package main

import (
	_ "github.com/Alfinaditya/go-fiber-simple-rest-api/docs"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/server"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"
	"log"

	"github.com/joho/godotenv"
)

// @title Fiber Example API
// @version 1.0
// @description API for managing books, authors, and users with JWT authentication

// @host localhost:3001
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database.ConnectDB()

	server.StartServer()
}
