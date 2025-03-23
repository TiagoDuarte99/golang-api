package dto

// LoginData estrutura para representar os dados de login.
type LoginRequest struct {
	Email    string `json:"email" example:"tiago@example.com" validate:"required,email"`
	Password string `json:"password" example:"SenhaForte123!" validate:"required,min=6"`
}
