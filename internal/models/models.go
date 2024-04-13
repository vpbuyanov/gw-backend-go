package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	HashPass string `json:"hash_pass"`
}
