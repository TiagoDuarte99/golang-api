package models

import (
	"time"
)

type User struct {
	ID            int       `gorm:"primary_key;autoIncrement" json:"id" db:"id"`
	Name          string    `gorm:"type:varchar(100);not null" json:"name" db:"name" validate:"required,min=2,max=100"`
	Email         string    `gorm:"type:varchar(100);unique;not null" json:"email" db:"email" validate:"email,required"`
	Password      *string   `gorm:"type:varchar(255);not null" json:"password" db:"password" validate:"required,min=6"`
	UserType      *string   `gorm:"type:varchar(20);not null" json:"user_type" db:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token         *string   `gorm:"type:varchar(255)" json:"token" db:"token"`
	RefreshToken  *string   `gorm:"type:varchar(255)" json:"refresh_token" db:"refresh_token"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at" db:"created_at"`
}
