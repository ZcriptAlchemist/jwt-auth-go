package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// // ----------------------------------------------------------------------
// // Generates random salt value to be padded with password before hashing
// // ----------------------------------------------------------------------
// func GenerateSalt() string {
// 	// Creating a new byte slice
// 	salt := make([]byte, 16)

// 	// Populating salt slice with random bytes
// 	_, err := rand.Read(salt)

// 	if err != nil {
// 		 advice.NewError( "error while generating salt")
// 	}

// 	// Returning the byte slice as string
// 	return base64.StdEncoding.EncodeToString(salt)
// }

// -----------------------------------------------------------------------------
// HashPassword adds salt to plain password and then hashes the salted password.
// -----------------------------------------------------------------------------

// func HashPassword(password string, salt string) (string, error){
// my function
// func HashPassword(password string) (string, error){

// 	// Hashing saltedPassed using bcrypt algorithm
//     hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	if err != nil {
//         return "error", err
//     }

// 	return string(hash), nil
// }

// ram bro function
// ----------------------------------------------------------------------------
// HashPassword takes a plaintext password and returns a bcrypt hashed version.
// ----------------------------------------------------------------------------
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// ----------------------------------------------------------------------------
// ComparePassword checks if a plaintext password matches a hashed password.
// ----------------------------------------------------------------------------
func ComparePassword(hashedPassword, plainPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	log.Println("from utils",err)
	return err != nil
}
 


// func ComparePasswords (loginPassword string, dbPassword string) error  {
// 	err := bcrypt.CompareHashAndPassword([]byte(loginPassword), []byte(dbPassword))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }