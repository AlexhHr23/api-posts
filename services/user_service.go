package services

import (
	"context"
	"fmt"

	"github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(ctx context.Context, name, email, password string) (*models.User, error) {
	exist, err := s.repo.EmailExist(ctx, email)

	if err != nil {
		return nil, err
	}

	if exist {
		return nil, fmt.Errorf("El email ya esta registrado")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("Erro al hashear la contraseña: %w", err)
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
