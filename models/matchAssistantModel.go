package models

type MatchAssistant struct {
	MatchID    int `gorm:"primaryKey;not null;index;foreignKey:MatchID;references:matches.id" json:"match_id"`
	PlayerID      int    `gorm:"primary_key;not null;index;foreignKey:ID;references:players.id" json:"player_id"` // Chave estrangeira para Player
	AssistedGoals int    `json:"assisted_goals"`
}
