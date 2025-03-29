package dto

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,len=12"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}
