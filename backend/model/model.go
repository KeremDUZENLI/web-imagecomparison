package model

import (
	"time"
)

type ProjectModel struct {
	ID                int       `json:"id,omitempty"`
	UserName          string    `json:"userName"`
	ImageA            string    `json:"imageA"`
	ImageB            string    `json:"imageB"`
	ImageWinner       string    `json:"imageWinner"`
	ImageLoser        string    `json:"imageLoser"`
	EloWinnerPrevious int       `json:"eloWinnerPrevious"`
	EloWinnerNew      int       `json:"eloWinnerNew"`
	EloLoserPrevious  int       `json:"eloLoserPrevious"`
	EloLoserNew       int       `json:"eloLoserNew"`
	CreatedAt         time.Time `json:"createdAt,omitempty"`
}
