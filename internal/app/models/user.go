package models

// User model
type User struct {
	ID             int    `json:"id"`
	Login          string `json:"login"`
	HashedPassword string `json:"hashedpassword"`
	SecretNumber   string `json:"secretnumber"`
}
