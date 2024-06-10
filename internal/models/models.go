package models

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	HashPass string `json:"hash_pass"`
	IsAdmin  bool   `json:"is_admin"`
	IsBanned bool   `json:"is_banned"`
}
