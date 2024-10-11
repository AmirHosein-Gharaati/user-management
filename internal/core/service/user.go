package service

import (
	"errors"
	"log/slog"

	"github.com/AmirHosein-Gharaati/user-management/internal/core/domain"
	"github.com/AmirHosein-Gharaati/user-management/internal/core/port"
)

type UserServiceImpl struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Register(user *domain.User) (*domain.User, error) {
	exists := s.userRepo.ExistsUserByEmail(user.Email)
	if exists {
		slog.Error("user exists by the email")
		return nil, errors.New("user exists by the email")
	}

	userDB, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return userDB, nil
}
