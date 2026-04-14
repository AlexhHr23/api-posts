package middleware

import (
	"net/http"
	"strings"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/handlers"
	"github.com/AlexhHr23/gopost-api/server"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next server.HandleFunc) server.HandleFunc {
	return func(c *server.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			handlers.RespondError(c, handlers.NewAppError("Token no proporcionado", http.StatusUnauthorized))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			handlers.RespondError(c, handlers.NewAppError("Formatio de token invalido", http.StatusUnauthorized))
			return
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, handlers.NewAppError("Metodo de firma inesperado", http.StatusUnauthorized)
			}

			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			handlers.RespondError(c, handlers.NewAppError("claims invalido", http.StatusUnauthorized))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			handlers.RespondError(c, handlers.NewAppError("Token invalido", http.StatusUnauthorized))
			return
		}

		userId, ok := claims["user_id"].(float64)

		if !ok {
			handlers.RespondError(c, handlers.NewAppError("User ID no encontrado en el token", http.StatusUnauthorized))
			return
		}

		c.SetUserID(uint(userId))
		next(c)
	}
}
