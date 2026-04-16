package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/server"
	"github.com/AlexhHr23/gopost-api/services"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatetPost(c *server.Context) {
	var req models.Post

	if err := c.BindJson(&req); err != nil {
		RespondError(c, NewAppError("Datos invalidos", http.StatusBadRequest))
		return
	}

	if req.Title == "" || req.Content == "" {
		RespondError(c, NewAppError("Todos los campos son obligatorios", http.StatusBadRequest))
		return
	}

	post, err := h.postService.CreatePost(c.Context(), c.UserID, req.Title, req.Content)

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Post creado exitosamente",
		"post": map[string]interface{}{
			"id":         post.ID,
			"title":      post.Title,
			"content":    post.Content,
			"user_id":    post.UserID,
			"created_at": post.CreatedAt,
			"updated_at": post.UpdatedAt,
		},
	})
}

func (h *PostHandler) UpdatetPost(c *server.Context) {

	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req models.Post

	if err := c.BindJson(&req); err != nil {
		RespondError(c, NewAppError("Datos invalidos", http.StatusBadRequest))
		return
	}

	if req.Title == "" || req.Content == "" {
		RespondError(c, NewAppError("Todos los campos son obligatorios", http.StatusBadRequest))
		return
	}

	err := h.postService.UpdatePost(c.Context(), req.Title, req.Content, uint(id), c.UserID)

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Post actualizado exitosamente",
	})
}

func (h *PostHandler) GetPosts(c *server.Context) {
	posts, err := h.postService.GetAllPost(c.Context())

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "ok",
		"posts":  posts,
	})
}

func (h *PostHandler) DeletePost(c *server.Context) {

	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	err := h.postService.DeletePost(c.Context(), uint(id), c.UserID)

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Post eliminado exitosamente",
	})
}

//TODO
// func GetPostById(c *server.Context) {
// 	idStr := c.Request.PathValue("id")
// 	id, _ := strconv.Atoi(idStr)
// 	for _, post := range posts {
// 		if post.Id == id {
// 			c.JSON(http.StatusOK, post)
// 			return
// 		}
// 	}

// 	http.Error(c.RWriter, "Post no encontrado", http.StatusNotFound)
// }
