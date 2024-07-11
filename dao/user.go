package dao

import (
	"jwt-auth-go/api/advice"
	"jwt-auth-go/config"
	"jwt-auth-go/models"
	"net/http"

	"gorm.io/gorm"
)

// ----------------------------------
// Saves user data into the Database.
// ----------------------------------
func SaveUser(user *models.User) error {
	err := config.DB.Create(&user).Error

	if err != nil {
		return advice.NewAPIError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// ------------------------------------------
// Finds and returns users from the DataBase.
// ------------------------------------------
func FindUsers() ([]models.User, error) {

	var users []models.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		return users, advice.NewAPIError(http.StatusInternalServerError, result.Error.Error())
	}

	return users, nil
}

// ----------------------------------------------------------------
// Removes all users from the DB - this method is only for testing
// ----------------------------------------------------------------
func RemoveAllUsers () error {
	// DB.Exec is used to execute raw SQL queries
	err := config.DB.Exec("DELETE FROM users")

	if err != nil {
		return err.Error
	}

	return nil
}

// ------------------------------------------------------
// Finds the user in the DB based on entered credentials.
// ------------------------------------------------------
// func FindUserByEmail (email *string) (*models.User, error) {
// 	var user models.User

// 	result := config.DB.First(&user, "email = ?", email)

// 	if result.Error != nil {
// 	if result.Error == gorm.ErrRecordNotFound {
// 		return nil, advice.NewAPIError(http.StatusNotFound, "User not found") // Proper handling for user not found
// 	}
// 	return nil, result.Error // Ensured all errors are correctly returned
// }


// 	if user.ID == 0 {
// 		return nil, advice.NewAPIError(http.StatusNotFound, result.Error.Error())
// 	}

// 	return &user, nil
// }

func FindUserByEmail(email *string) (*models.User, error) {
	var user models.User

	result := config.DB.Table("users").Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, advice.NewAPIError(http.StatusNotFound, "User not found")
		}
		return nil, result.Error
	}

	return &user, nil
}
