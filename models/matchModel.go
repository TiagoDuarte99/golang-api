package models

import (
	"time"
)

type Match struct {
	ID         int       `gorm:"primary_key;autoIncrement" json:"id"`
	HomeTeamID int       `gorm:"not null;index;foreignKey:ID;references:teams.id" json:"home_team_id"`
	AwayTeamID int       `gorm:"not null;index;foreignKey:ID;references:teams.id" json:"away_team_id"`
	Date       time.Time `json:"date"`
	Location   string    `gorm:"type:varchar(100);not null" json:"location"`
	HomeGoals  int       `json:"home_goals"`
	AwayGoals  int       `json:"away_goals"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	Scorers    []MatchScorer   `gorm:"foreignKey:MatchID" json:"scorers"` // Relacionamento de 1 para muitos
	Assistants []MatchAssistant `gorm:"foreignKey:MatchID" json:"assistants"` // Relacionamento de 1 para muitos
}
