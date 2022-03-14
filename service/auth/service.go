package auth

import (
    "github.com/templateOfService/model"
    authRepo "github.com/templateOfService/repo/auth"
)

type Service struct {
    repo *authRepo.Repo
}

func NewService() *Service {
    return &Service{
        repo: authRepo.NewRepo(),
    }
}

func (s *Service) Signup(user model.User) string {
    return "signed up"
}
