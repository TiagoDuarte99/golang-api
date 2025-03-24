package dto

import (
	"time"
)

type UserResponse struct {
	ID           int       `json:"id" example:"1"`
	Name         string    `json:"name" example:"Tiago Duarte"`
	Email        string    `json:"email" example:"tiago@example.com"`
	UserType     string    `json:"user_type" example:"USER"`
	CreatedAt    time.Time `json:"created_at" example:"2024-03-24T15:30:00Z"`
}


type UpdateUserRequest struct {
	Name            string `json:"name" example:"John Doe"`
	Email           string `json:"email" example:"johndoe@example.com"`
	Password        string `json:"password" example:"newpassword123"`
	ConfirmPassword string `json:"confirm_password" example:"newpassword123"`
}
