package app

import (
	"encoding/json"
	"strings"
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

type RatingsModel struct {
	Image string `json:"image"`
	Elo   int    `json:"elo"`
}

type VotesDTO struct {
	Username    string `json:"username"`
	ImageWinner string `json:"imageWinner"`
	ImageLoser  string `json:"imageLoser"`
}

func (v *VotesDTO) UnmarshalJSON(raw []byte) error {
	type alias VotesDTO
	aux := struct{ *alias }{alias: (*alias)(v)}

	if err := json.Unmarshal(raw, &aux); err != nil {
		return err
	}

	v.Username = strings.ToLower(strings.TrimSpace(v.Username))
	return nil
}
