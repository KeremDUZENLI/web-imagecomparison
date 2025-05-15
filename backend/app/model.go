package app

import (
	"time"
)

type VotesModel struct {
	Username          string    `json:"username"`
	ImageWinner       string    `json:"imageWinner"`
	ImageLoser        string    `json:"imageLoser"`
	EloWinnerPrevious int       `json:"eloWinnerPrevious"`
	EloWinnerNew      int       `json:"eloWinnerNew"`
	EloLoserPrevious  int       `json:"eloLoserPrevious"`
	EloLoserNew       int       `json:"eloLoserNew"`
	CreatedAt         time.Time `json:"createdAt,omitempty"`
}

type VotesDTO struct {
	Username    string `json:"username"`
	ImageWinner string `json:"imageWinner"`
	ImageLoser  string `json:"imageLoser"`
}

type RatingsModel struct {
	Image string `json:"image"`
	Elo   int    `json:"elo"`
}
