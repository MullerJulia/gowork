package users

import (
	"context"
	"fmt"

	"github.com/mikenai/gowork/internal/models"
)

//go:generate
type Repositry interface {
	Create(ctx context.Context, name, phoneNumber string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
}

type Service struct {
	repo Repositry
}

func New(r Repositry) Service {
	return Service{repo: r}
}

func (s Service) Create(ctx context.Context, name, phoneNumber string) (models.User, error) {
	if name == "" {
		return models.User{}, fmt.Errorf("invalid name argument: %w", models.UserCreateParamInvalidNameErr)
	}

	usr, err := s.repo.Create(ctx, name, phoneNumber)
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

func (s Service) UpdateUser(ctx context.Context, id, name, phoneNumber string) (models.User, error) {
	if id == "" || name == "" {
		return models.User{}, fmt.Errorf("one of the required parameters is empty: %w", models.InvalidErr)
	}

	user := models.User{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
	}

	err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return user, fmt.Errorf("failed to update user's name: %w", err)
	}

	return user, nil
}
