package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	accessToken := context.Request.Header.Get("Authorization")
	
	if accessToken == "" {
		context.JSON(http.StatusUnauthorized,"Missing token")
		context.Abort()
		return
	}

	strings.Split(accessToken, "Bearer ")



	context.Set("useremail" , "")
	context.Set("user role", "user")

}