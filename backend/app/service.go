package app

import (
	"context"
	"math"
	"web-imagecomparison/env"
)

type ProjectService struct {
	Repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (ps *ProjectService) ProcessVote(ctx context.Context, dto *VotesDTO) (*VotesModel, error) {
	ratings, err := ps.Repo.GetAllTableRatings(ctx)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]int, len(ratings))
	for _, r := range ratings {
		lookup[r.Image] = r.Elo
	}

	prevW := lookup[dto.ImageWinner]
	if prevW == 0 {
		prevW = env.DEFAULT_RATING
	}
	prevL := lookup[dto.ImageLoser]
	if prevL == 0 {
		prevL = env.DEFAULT_RATING
	}

	ea := 1.0 / (1.0 + math.Pow(10, float64(prevL-prevW)/400.0))
	delta := int(math.Round(env.K_FACTOR * (1.0 - ea)))

	vote := &VotesModel{
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

	if err := ps.Repo.InsertTableRatings(
		ctx,
		RatingsModel{Image: vote.ImageWinner, Elo: vote.EloWinnerNew},
		RatingsModel{Image: vote.ImageLoser, Elo: vote.EloLoserNew},
	); err != nil {
		return nil, err
	}

	return vote, nil
}

func (ps *ProjectService) GetAllRatings(ctx context.Context) ([]RatingsModel, error) {
	return ps.Repo.GetAllTableRatings(ctx)
}
