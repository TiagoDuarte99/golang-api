package dto

type LoginRequest struct {
	Email    string `json:"email" example:"tiago@example.com" validate:"required,email"`
	Password string `json:"password" example:"SenhaForte123!" validate:"required,min=6"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	Token       string   `json:"token"`
	RefreshToken string  `json:"refresh_token"`
}

type SignupRequest struct {
	Name            string    `json:"name" example:"Tiago Duarte" validate:"required,min=2,max=100"`
	Email           string    `json:"email" example:"tiago@example.com" validate:"required,email"`
	Password        string    `json:"password" example:"SenhaForte123!" validate:"required,min=6"`
	ConfirmPassword string    `json:"confirm_password" example:"SenhaForte123!" validate:"required,min=6"`
	UserType        string    `json:"user_type" example:"USER" validate:"required,eq=USER"`
}