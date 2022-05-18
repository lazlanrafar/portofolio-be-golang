package loginAuth

import model "lazlanrafar/models"

type Service interface {
	LoginRepository(username string, password string) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository}
}

func (s *service) LoginRepository(username string, password string) (*model.User, string) {
	return s.repository.LoginRepository(username, password)
}