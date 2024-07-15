package services

import (
	"fmt"

	// "jwt-auth-go/api/advice"
	"jwt-auth-go/dao"
	"jwt-auth-go/dto"
	"jwt-auth-go/models"
	"jwt-auth-go/tokens"
	"jwt-auth-go/utils"
)

// -------------
// creates user.
// -------------
func CreateUser(user *models.User) error {
	// Mapping generated salt value to Salt field in user struct
	user.Salt = utils.GenerateSalt()

	userPass := user.Password + user.Salt

	// Hashing the password.
	hashedPassword, err := utils.HashPassword(userPass)

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
func FetchLoginCredentials(credentials *dto.Credentials)  (*dto.LoginResponse, error) {

	user, err := dao.FindUserByEmail(&credentials.Email)
	if err != nil {
		return nil,err
	}

	// Ensure user is not nil before proceeding
	if user == nil {
		return nil,fmt.Errorf("user not found")
	}

	// Adding salt fetched from Db to login password.
	plain := credentials.Password + user.Salt

	// Compare sent in password with DB saved password
	err = utils.ComparePassword(user.Password, plain)

	if err != nil {
		return  nil,err
	}

	// generate JWT Token
	token, err := tokens.GenerateJwtToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
	} , nil 

}