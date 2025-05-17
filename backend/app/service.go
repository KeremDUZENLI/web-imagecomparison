package app

import (
	"context"
	"math"
	"web-imagecomparison/env"
)

type ProjectService interface {
	GetAllUsernames(ctx context.Context) ([]string, error)
	GetAllRatings(ctx context.Context) ([]RatingsModel, error)
	PostSurvey(ctx context.Context, surveysModel SurveysModel) error
	PostVote(ctx context.Context, dto *VotesDTO) (*VotesModel, error)
}

type projectService struct {
	repository ProjectRepository
}

func NewProjectService(repo ProjectRepository) ProjectService {
	return &projectService{repository: repo}
}

func (ps *projectService) GetAllUsernames(ctx context.Context) ([]string, error) {
	return ps.repository.GetUsernames(ctx)
}

func (ps *projectService) GetAllRatings(ctx context.Context) ([]RatingsModel, error) {
	return ps.repository.GetRatings(ctx)
}

func (ps *projectService) PostSurvey(ctx context.Context, surveysModel SurveysModel) error {
	return ps.repository.InsertSurvey(ctx, surveysModel)
}

func (ps *projectService) PostVote(ctx context.Context, dto *VotesDTO) (*VotesModel, error) {
	ratings, err := ps.repository.GetRatings(ctx)
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
		Username:          dto.Username,
		ImageWinner:       dto.ImageWinner,
		ImageLoser:        dto.ImageLoser,
		EloWinnerPrevious: prevW,
		EloWinnerNew:      prevW + delta,
		EloLoserPrevious:  prevL,
		EloLoserNew:       prevL - delta,
	}

	if err := ps.repository.InsertVote(ctx, vote); err != nil {
		return nil, err
	}

	if err := ps.repository.InsertRating(
		ctx,
		RatingsModel{Image: vote.ImageWinner, Elo: vote.EloWinnerNew},
		RatingsModel{Image: vote.ImageLoser, Elo: vote.EloLoserNew},
	); err != nil {
		return nil, err
	}

	return vote, nil
}
