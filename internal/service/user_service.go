package service

import (
	"context"

	"github.com/nishant007-tech/GoDig/internal/logger"
	"github.com/nishant007-tech/GoDig/internal/repository"
	"go.uber.org/dig"
)

// UserService defines business methods.
type UserService interface {
	List(ctx context.Context) ([]repository.User, error)
}

// Params for NewUserService.
type userServiceParams struct {
	dig.In

	Repo   repository.UserRepository
	Logger *logger.Logger `optional:"true"`
}

// NewUserService constructs a service with its dependencies.
func NewUserService(p userServiceParams) UserService {
	return &userService{repo: p.Repo, log: p.Logger}
}

type userService struct {
	repo repository.UserRepository
	log  *logger.Logger
}

func (s *userService) List(ctx context.Context) ([]repository.User, error) {
	if s.log != nil {
		s.log.Info("service: listing users")
	}
	return s.repo.GetAll(ctx)
}
