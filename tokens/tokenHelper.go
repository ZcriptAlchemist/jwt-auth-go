package tokens

import (
	"jwt-auth-go/api/advice"
	"jwt-auth-go/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// ---------------------
// Generating JWT token.
// ---------------------
func GenerateJwtToken(user *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    user.ID,
		"userEmail": user.Email,
		"exp":       time.Now().Add(time.Minute * 20).Unix(),
	})

	hmacSampleSecret := os.Getenv("JWT_SECRET")

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	if err != nil {
		return "", advice.NewError(err.Error())
	}


	return tokenString, nil
}