package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlexhHr23/gopost-api/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

//Create post

func (r *PostRepository) Create(ctx context.Context, post *models.Post) error {
	query := "INSERT INTO posts(user_id, title, content) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, post.UserID, post.Title, post.Content)
	if err != nil {
		return fmt.Errorf("Error al crear post: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el ID: %w", err)
	}
	post.ID = uint(id)
	return nil
}

func (r *PostRepository) FindAll(ctx context.Context) ([]models.Post, error) {
	query := "SELECT * FROM posts ORDER BY createt_at DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener posts: %w", err)
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Error al escanear post: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) FindById(ctx context.Context, id uint) (*models.Post, error) {
	post := &models.Post{}
	query := "SELECT * FROM posts WHERE id = ?"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Post no encontrado: %w", err)
		}
		return nil, fmt.Errorf("Error al buscar post: %w", err)
	}
	return post, nil
}

func (r *PostRepository) FindByUserId(ctx context.Context, userId uint) ([]models.Post, error) {
	query := "SELECT * FROM posts WHERE user_id = ?"
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener posts del usuario: %w", err)
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Error al escanear post: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) Update(ctx context.Context, post *models.Post, id uint) error {
	query := "UPDATE posts SET title = ?, content = ? WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, post.Title, post.Content, id)

	if err != nil {
		return fmt.Errorf("Error al actualizar post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Error al verificar actualizacion: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Post no encontrado")
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM posts WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("error al eliminar postL: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Error al verificar eliminacion: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Post no encontrado")
	}

	return nil
}
