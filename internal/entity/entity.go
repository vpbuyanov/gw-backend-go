package entity

type LoginUserRequest struct {
	Email    string `json:"email"`
	HashPass string `json:"hash_pass"`
}

type RegistrationUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	HashPass string `json:"hash_pass"`
}
