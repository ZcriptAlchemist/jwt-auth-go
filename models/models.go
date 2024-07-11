package models

import "gorm.io/gorm"

// --------------------------------
// User struct to define DB Schema.
// --------------------------------
type User struct {
	gorm.Model
	Name string
	Email string
	Number string
	Password string
	// Salt string
}

// --------------------
// Credentials structs.
// --------------------
type Credentials struct {
	Email string
	Password string
}