package user

import (
	"errors"

	"github.com/julioshinoda/challenge-api/pkg/auth"
)

type Manager interface {
	Signin(user User) (string, error)
}

type service struct {
	Repo Repository
}

func NewUser() Manager {
	return service{Repo: NewRepo()}
}

func (s service) Signin(user User) (string, error) {
	persistedUser, err := s.Repo.GetUser(user.Username)
	if err != nil {
		return "", err
	}
	if persistedUser.Secret == user.Secret {
		return auth.GenerateToken(persistedUser.ID), nil
	}
	return "", errors.New("user not found")
}
