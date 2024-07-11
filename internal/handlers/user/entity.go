package user

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Phone    string `json:"phone"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
