package models

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	HashPass string `json:"hash_pass"`
	IsAdmin  bool   `json:"is_admin"`
}
