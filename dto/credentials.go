package dto

// --------------------
// Credentials structs.
// --------------------
type Credentials struct {
	Email string
	Password string
}


type LoginResponse struct {
	AccessToken string
}