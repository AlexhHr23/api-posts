package handlers

import (
	"net/http"

	"github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/server"
	"github.com/AlexhHr23/gopost-api/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUpHandler(c *server.Context) {
	var req models.SignUpInput
	if err := c.BindJson(&req); err != nil {
		RespondError(c, NewAppError("Datos invalidos", http.StatusBadRequest))
		return
	}

	if req.Name == "" || req.Email == "" || req.Passowrd == "" {
		RespondError(c, NewAppError("Todos los campos son obligatorios", http.StatusBadRequest))
		return
	}

	user, err := h.userService.SignUp(c.Context(), req.Name, req.Email, req.Passowrd)

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Usuario  creado exitosamente",
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
