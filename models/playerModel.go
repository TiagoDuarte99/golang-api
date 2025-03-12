package models

import (
	"time"
)

type Player struct {
	ID       int       `gorm:"primary_key;autoIncrement" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Position string    `gorm:"type:varchar(50);not null" json:"position"`
	TeamID   int       `gorm:"not null;index" json:"team_id"`
	Age      int       `json:"age"`
	Height   float64   `json:"height"`
	Weight   float64   `json:"weight"`
	Goals    int       `gorm:"default:0" json:"goals"`
	Assists  int       `gorm:"default:0" json:"assists"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	Team     Team      `gorm:"foreignKey:TeamID;references:id" json:"team"`
}
