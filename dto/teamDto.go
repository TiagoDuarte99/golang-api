package dto

type CoachDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TeamDTO struct {
	ID        int     `json:"id"`
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	CoachID   int     `json:"coach_id"`
	Pts      int      `json:"pts"`
	CreatedAt string   `json:"created_at"`
	Coach     CoachDTO `json:"coach"`
}