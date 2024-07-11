package main

import (
	"jwt-auth-go/config"
	"jwt-auth-go/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// -------------------------------------------------
// Execution starts here, Initializing a web server.
// -------------------------------------------------
func main() {
	// Initializing a web server.
	router := gin.New()
	
	// Connecting to DB.
	err := config.Connect()

	if err != nil {
		log.Println(err)
	}

	// Passing router or web server into user defined package of routes
	routes.UserRoutes(router)
	
	// Running the server 5000.
	router.Run(":5000")
}