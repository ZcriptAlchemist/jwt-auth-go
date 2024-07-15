package utils

import (
	"crypto/rand"
	"encoding/base64"
	"jwt-auth-go/api/advice"

	"golang.org/x/crypto/bcrypt"
)

// ----------------------------------------------------------------------
// Generates random salt value to be padded with password before hashing
// ----------------------------------------------------------------------
func GenerateSalt() string {
	// Creating a new byte slice
	salt := make([]byte, 16)

	// Populating salt slice with random bytes
	_, err := rand.Read(salt)

	if err != nil {
		 advice.NewError( "error while generating salt")
	}

	// Returning the byte slice as string
	return base64.StdEncoding.EncodeToString(salt)
}

// ----------------------------------------------------------------------------
// HashPassword takes a plaintext password and returns a bcrypt hashed version.
// ----------------------------------------------------------------------------
func HashPassword(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	
	if err != nil {
		return "", err
	}
	
	return string(bs), nil
}



// ----------------------------------------------------------------------------
// ComparePassword checks if a plaintext password matches a hashed password.
// ----------------------------------------------------------------------------
func ComparePassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}
	return nil
}