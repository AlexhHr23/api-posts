package services

import (
	"context"
	"fmt"
	"regexp"

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

// Validar email
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("Formato de email no valido")
	}

	return nil
}

func ValidatePasswotd(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("La contraseña debe tener al menos 8 caracteres")
	}
	return nil
}

func (s *UserService) SignUp(ctx context.Context, name, email, password string) (*models.User, error) {

	if err := ValidateEmail(email); err != nil {
		return nil, err
	}

	if err := ValidatePasswotd(password); err != nil {
		return nil, err
	}

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
