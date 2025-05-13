package app

type ProjectService struct {
	Repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) CreateServiceEntry(v *ProjectModel) error {
	if err := s.Repo.InsertTableVotes(v); err != nil {
		return err
	}
	delta := v.EloWinnerNew - v.EloWinnerPrevious
	return s.Repo.UpdateTableRatings(v.ImageWinner, v.ImageLoser, delta)
}

func (s *ProjectService) GetallServiceRatings() (map[string]int, error) {
	return s.Repo.GetallTableRatings()
}

func (s *ProjectService) UpdateServiceRatings(winner, loser string, delta int) error {
	return s.Repo.UpdateTableRatings(winner, loser, delta)
}
