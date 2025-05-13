package app

type ProjectService struct {
	Repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) RecordVote(v *ProjectModel) error {
	return s.Repo.InsertTableVotes(v)
}

func (s *ProjectService) RecalculateRatings(v *ProjectModel) error {
	delta := v.EloWinnerNew - v.EloWinnerPrevious
	return s.Repo.UpdateTableRatings(v.ImageWinner, v.ImageLoser, delta)
}

func (s *ProjectService) GetAllRatings() (map[string]int, error) {
	return s.Repo.GetAllTableRatings()
}

func (s *ProjectService) UpdateRatings(winner, loser string, delta int) error {
	return s.Repo.UpdateTableRatings(winner, loser, delta)
}
