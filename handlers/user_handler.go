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

func (h *UserHandler) LoginHandler(c *server.Context) {
	var req models.LoginIput

	if err := c.BindJson(&req); err != nil {
		RespondError(c, NewAppError("Datos invalidos", http.StatusBadRequest))
		return
	}

	if req.Email == "" || req.Passowrd == "" {
		RespondError(c, NewAppError("Email y contraseña son obligatoriso", http.StatusBadRequest))
	}

	token, err := h.userService.Login(c.Context(), req.Email, req.Passowrd)

	if err != nil {
		RespondError(c, NewAppError(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Incio de sesion exitoso",
		"token":   token,
	})
}

func (h *UserHandler) MeHandler(c *server.Context) {
	userID := c.UserID
	if userID == 0 {
		RespondError(c, NewAppError("Usuario no autenticado", http.StatusFound))
		return
	}

	user, err := h.userService.GetUserByID(c.Context(), userID)

	if err != nil {
		RespondError(c, NewAppError("Usuario no encontrado", http.StatusFound))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
