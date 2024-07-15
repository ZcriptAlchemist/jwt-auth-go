package routes

import (
	"jwt-auth-go/api/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	// Create user.
	router.POST("/users/create", controller.CreateUser)
	
	// Login route.
	router.POST("users/authenticate", controller.Login)
}

func UserRoutes(router gin.IRoutes)  {

	// Get all users.
	router.GET("/get-users", controller.GetUsers)
	
	// Deletes all users.
	router.DELETE("/delete-users", controller.DeleteAllUsers)
}