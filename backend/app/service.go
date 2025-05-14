package app

import (
	"context"
	"math"
	"time"
	"web-imagecomparison/env"
)

type ProjectService struct {
	Repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (ps *ProjectService) ProcessVote(ctx context.Context, dto *VoteDTO) (*VoteModel, error) {
	ratings, err := ps.Repo.GetAllTableRatings(ctx)
	if err != nil {
		return nil, err
	}

	prevW, ok := ratings[dto.ImageWinner]
	if !ok {
		prevW = env.DEFAULT_RATING
	}
	prevL, ok := ratings[dto.ImageLoser]
	if !ok {
		prevL = env.DEFAULT_RATING
	}

	ea := 1.0 / (1.0 + math.Pow(10, float64(prevL-prevW)/400.0))
	delta := int(math.Round(env.K_FACTOR * (1.0 - ea)))

	vote := &VoteModel{
		UserName:          dto.UserName,
		ImageA:            dto.ImageA,
		ImageB:            dto.ImageB,
		ImageWinner:       dto.ImageWinner,
		ImageLoser:        dto.ImageLoser,
		EloWinnerPrevious: prevW,
		EloWinnerNew:      prevW + delta,
		EloLoserPrevious:  prevL,
		EloLoserNew:       prevL - delta,
	}

	if err := ps.Repo.InsertTableVotes(ctx, vote); err != nil {
		return nil, err
	}

	if vote.CreatedAt.IsZero() {
		vote.CreatedAt = time.Now()
	}

	if err := ps.Repo.UpdateTableRatings(ctx, vote.ImageWinner, vote.ImageLoser, delta); err != nil {
		return nil, err
	}

	return vote, nil
}

func (ps *ProjectService) GetAllRatings(ctx context.Context) (map[string]int, error) {
	return ps.Repo.GetAllTableRatings(ctx)
}
