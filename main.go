package main

import (
	"jwt-auth-go/config"
	"jwt-auth-go/middleware"
	"jwt-auth-go/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// -------------------------------------------------
// Execution starts here, Initializing a web server.
// -------------------------------------------------
func main() {
	// Loading environment variables.
	godotenv.Load()
	// Initializing a web server.
	router := gin.New()
	
	// Connecting to DB.
	err := config.Connect()

	if err != nil {
		log.Println(err)
	}

	// Passing router or web server into user defined package of routes
	routes.AuthRoutes(router)

	// Group routes that require authentication
	protected := router.Group("/")
	protected.Use(middleware.Authenticate)
	routes.UserRoutes(protected)
	
	// Running the server 5000.
	router.Run(":5000")
}