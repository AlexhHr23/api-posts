package services

import (
	"context"
	"fmt"

	"github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

// Create post
func (s *PostService) CreatePost(ctx context.Context, userId uint, title, content string) (*models.Post, error) {
	post := &models.Post{
		UserID:  userId,
		Title:   title,
		Content: content,
	}

	if err := s.repo.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

// Update post
func (s *PostService) UpdatePost(ctx context.Context, title, content string, id, userID uint) error {

	findPost, err := s.repo.FindById(ctx, id)

	if err != nil {
		return err
	}

	if userID != findPost.UserID {
		return fmt.Errorf("Solo el usario que creo el post puede editarlo")
	}

	post := &models.Post{
		Title:   title,
		Content: content,
	}

	if err := s.repo.Update(ctx, post, id); err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetAllPost(ctx context.Context) ([]models.Post, error) {
	return s.repo.FindAll(ctx)
}

func (s *PostService) DeletePost(ctx context.Context, id, userID uint) error {
	findPost, err := s.repo.FindById(ctx, id)

	if err != nil {
		return err
	}

	if userID != findPost.UserID {
		return fmt.Errorf("Solo el usario que creo el post puede eliminarlo")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
