package app

import (
	"time"
)

type SurveysModel struct {
	Username   string    `json:"username"`
	Age        string    `json:"age"`
	Gender     string    `json:"gender"`
	Experience string    `json:"experience"`
	Profession string    `json:"profession"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type VotesModel struct {
	Username          string    `json:"username"`
	ImageWinner       string    `json:"image_winner"`
	ImageLoser        string    `json:"image_loser"`
	EloWinnerPrevious int       `json:"elo_winner_previous"`
	EloWinnerNew      int       `json:"elo_winner_new"`
	EloLoserPrevious  int       `json:"elo_loser_previous"`
	EloLoserNew       int       `json:"elo_loser_new"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

type RatingsModel struct {
	Image     string    `json:"image"`
	Elo       int       `json:"elo"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VotesDTO struct {
	Username    string `json:"username"`
	ImageWinner string `json:"image_winner"`
	ImageLoser  string `json:"image_loser"`
}
