package models

import (
	"time"
)

type Team struct {
	ID        int       `gorm:"primary_key;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Country   string    `gorm:"type:varchar(50);not null" json:"country"`
	CoachID   int       `gorm:"not null;index;foreignKey:CoachID;references:id" json:"coach_id"`
	Pts       int       `gorm:"default:0" json:"pts"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	Coach     User      `gorm:"foreignKey:CoachID;references:id" json:"coach"`
}
