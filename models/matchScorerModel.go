package models

type MatchScorer struct {
	MatchID    int `gorm:"primaryKey;not null;index;foreignKey:MatchID;references:matches.id" json:"match_id"`
	PlayerID   int `gorm:"primaryKey;not null;index;foreignKey:PlayerID;references:players.id" json:"player_id"`
	ScoredGoals int `json:"scored_goals"`
}
