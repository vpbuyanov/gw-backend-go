package entity

type LoginUserRequest struct {
	Email    string `json:"email"`
	HashPass string `json:"hash_pass"`
}

type RegistrationUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Phone    string `json:"phone"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
