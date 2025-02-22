package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"  validate:"required,min=2,max=100"`
	Email         string    `json:"email" db:"email" validate:"email,required"`
	Password      *string   `json:"Password" validate:"required,min=6"`
	User_type     *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
