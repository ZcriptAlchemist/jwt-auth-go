package services

import (
	"fmt"
	"jwt-auth-go/api/advice"
	"jwt-auth-go/dao"
	"jwt-auth-go/models"
	"jwt-auth-go/utils"
	"log"
)

// -------------
// creates user.
// -------------
func CreateUser(user *models.User) error {
	// Mapping generated salt value to Salt field in user struct
	// user.Salt = utils.GenerateSalt()

	// Hashing the password.
	// hashedPassword, err := utils.HashPassword(user.Password, user.Salt)
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}
	// Mapping hashedPassword to password field in user struct
	user.Password = hashedPassword

	err = dao.SaveUser(user) 

	if err != nil {
		return err
	}

	return nil
}

// ------------------
// Fetches user data
// ------------------
func FetchUsers () ([]models.User,error) {

	result, err := dao.FindUsers()

	if err != nil {
		return result, err
	}

	return result, nil
}


// ----------------------------------------------------------------
// Deletes all users from the DB - this method is only for testing
// ----------------------------------------------------------------
func DeleteAllUsers () error {
	
	err := dao.RemoveAllUsers()

	if err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------------------------
// Finds credentials in the DB and compares with user entered credentials for log in
// ----------------------------------------------------------------------------------
// func FetchLoginCredentials (credentials *models.Credentials) (*models.User, error) {

	
// 	user, err := dao.FindUserByEmail(&credentials.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// compare sent in password with DB saved password
// 	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
// 		if err2 != nil {
// 		return nil, err2
// 	}

// 	// // jwt sht
// 	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	// 	"sub": user.ID,
// 	// 	"exp": time.Now().Add(time.Minute * 15).Unix(),
// 	// })

// 	// // Sign and get the complete encoded token as a string using the secret
// 	// tokenString, err := token.SignedString(hmacSampleSecret)

// 	// fmt.Println(tokenString, err)

// 	return user, nil
// }

func FetchLoginCredentials(credentials *models.Credentials) (*models.User, error) {

	user, err := dao.FindUserByEmail(&credentials.Email)
	if err != nil {
		return nil, err
	}

	// Ensure user is not nil before proceeding
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Compare sent in password with DB saved password
	// signInPass, err3 := utils.HashPassword(credentials.Password, user.Salt)
	signInPass, err3 := utils.HashPassword(credentials.Password)
	if err3 != nil {
		log.Println(err3)
	}

	log.Println("password from db: ", user.Password)
	log.Println("password from entered password: ", signInPass)

	// err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	isMatch := utils.ComparePassword(user.Password, signInPass)

	if isMatch {
		return nil, advice.NewError("no match")
	}

	// Token generation and additional handling can be done here if needed

	return user, nil
}