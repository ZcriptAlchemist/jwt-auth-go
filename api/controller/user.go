package controller

import (
	"jwt-auth-go/models"
	"jwt-auth-go/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------------
// Creates user.
// -------------
func CreateUser(context *gin.Context) {
	
	var user models.User

	err := context.BindJSON(&user) 

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})		
	}

	err2 := services.CreateUser(&user)

	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err})
		return
	}

	context.JSON(http.StatusCreated,"successfully created the user")

}

// -----------------------
// Gets users from the DB.
// -----------------------
func GetUsers (context *gin.Context) {
	
	response, err := services.FetchUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err})
		return
	}

	context.JSON(http.StatusOK ,response)
}


// ----------------------------------------------------------------
// Deletes all users from the DB - this method is only for testing
// ----------------------------------------------------------------
func DeleteAllUsers(context *gin.Context) {

	err := services.DeleteAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err})
		return
	}
	
	context.JSON(http.StatusNoContent,"successfully deleted all users")
}

// --------
// Logs in
// --------
// func Login (context *gin.Context) {

// 	var credentials models.Credentials
	
// 	err := context.BindJSON(&credentials) 

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
// 		return
// 	}

// 	response, err := services.FetchLoginCredentials(&credentials)

// 	if err != nil {
// 		log.Println(response.Salt)
// 		log.Println("controller layer err: ", err)
// 		context.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
// 		return
// 	}
	
// 	context.JSON(http.StatusOK,response)
// 	log.Println(response.Email)
// }

func Login(context *gin.Context) {
	var credentials models.Credentials

	// Bind JSON and handle errors properly
	err := context.BindJSON(&credentials)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Fetch login credentials and handle errors
	response, err := services.FetchLoginCredentials(&credentials)
	if err != nil {
		// log.Println(response.Salt)
		log.Println("Error fetching login credentials:", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Ensure response is not nil before accessing fields
	if response == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error: user not found"})
		return
	}

	context.JSON(http.StatusOK, response)
	log.Println("Successful login for user:", response.Email)
}
