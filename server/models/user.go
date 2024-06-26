package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DisplayEmail string `json:"display_email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Error struct {
	message []byte `json:"message"`
}
