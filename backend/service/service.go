package service

import (
	"web-imagecomparison/model"
	"web-imagecomparison/repository"
)

type ProjectService struct {
	Repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) CreateVote(v *model.ProjectModel) error {
	return s.Repo.InsertVote(v)
}
