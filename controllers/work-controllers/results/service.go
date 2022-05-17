package resultsWork

import model "lazlanrafar/models"

type Service interface {
	ResultsWorkRepository() (*[]model.EntityWork, string)
}

type service struct{
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository}
}

func (s *service) ResultsWorkRepository() (*[]model.EntityWork, string) {
	return s.repository.ResultsWorkRepository()
}