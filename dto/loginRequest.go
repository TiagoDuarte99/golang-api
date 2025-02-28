package dto

// LoginData estrutura para representar os dados de login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}