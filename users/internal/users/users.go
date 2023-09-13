package users

import (
	"context"
	"fmt"

	"github.com/mikenai/gowork/internal/models"
)

type Repositry interface {
	Create(ctx context.Context, name string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	UpdateUserName(ctx context.Context, user models.User) error
}

type Service struct {
	repo Repositry
}

func New(r Repositry) Service {
	return Service{repo: r}
}

func (s Service) Create(ctx context.Context, name string) (models.User, error) {
	if name == "" {
		return models.User{}, fmt.Errorf("invalid name argument: %w", models.UserCreateParamInvalidNameErr)
	}

	usr, err := s.repo.Create(ctx, name)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return usr, nil
}

func (s Service) GetOne(ctx context.Context, id string) (models.User, error) {
	if id == "" {
		return models.User{}, fmt.Errorf("id is empty: %w", models.InvalidErr)
	}

	usr, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return usr, nil
}

func (s Service) UpdateUserName(ctx context.Context, user models.User) error {
	if user.ID == "" || user.Name == "" {
		return fmt.Errorf("one of user's param is empty: %w", models.InvalidErr)
	}

	err := s.repo.UpdateUserName(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user's name: %w", err)
	}

	return nil
}
