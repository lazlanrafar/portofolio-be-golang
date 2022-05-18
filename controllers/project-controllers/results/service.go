package resultsProject

import model "lazlanrafar/models"

type Service interface {
	ResultsProjectRepository() (*[]model.Project, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository}
}

func (s *service) ResultsProjectRepository() (*[]model.Project, string) {
	return s.repository.ResultsProjectRepository()
}