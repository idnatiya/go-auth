package main

import (
	"idnatiya/go-auth/databases"
	"idnatiya/go-auth/exceptions"
	"idnatiya/go-auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environtment variable
	godotenv.Load(".env")
	// Connect to Database
	databases.ConnectMysqlDB()
	// Create Main Instance Application
	route := gin.Default()
	// Set Global Exception Handler
	route.Use(exceptions.GlobalExceptionHandler())
	// Define Application Routes
	routes.DefineWebRoute(route)

	route.Run()
}
