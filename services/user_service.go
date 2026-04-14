package services

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/repositories"
	"github.com/golang-jwt/jwt/v5"
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

//Generar toke jwt

func (s *UserService) generateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// Logiun placeholder
func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		return "", fmt.Errorf("Credenciales incorrectas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", fmt.Errorf("Crendeciales incorrectas")
	}

	token, err := s.generateToken(user.ID)

	if err != nil {
		return "", fmt.Errorf("error al generar le token: %w", err)
	}

	return token, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}
