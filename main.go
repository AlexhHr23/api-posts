package main

import (

	// "fmt"
	"log"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/database"
	"github.com/AlexhHr23/gopost-api/handlers"
	"github.com/AlexhHr23/gopost-api/repositories"
	"github.com/AlexhHr23/gopost-api/services"

	// "github.com/AlexhHr23/gopost-api/models"

	"github.com/AlexhHr23/gopost-api/server"
)

// func hola(w http.ResponseWriter, r *http.Request) {

// 	name := r.URL.Query().Get("name")
// 	age := r.URL.Query().Get("age")

// 	datos := map[string]string{
// 		"name":    name,
// 		"age":     age,
// 		"mensaje": "Hola desde go",
// 		"status":  "OK",
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	err := json.NewEncoder(w).Encode(datos)

// 	if err != nil {
// 		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
// 	}
// }

func health(c *server.Context) {
	c.Send("Servidor corriendo")
}

func main() {

	config := config.LoadConfig()

	if err := database.Connect(config.DatabaseURL); err != nil {
		log.Fatal("Error al conector a la base de datos: ", err)
	}

	defer database.Close()

	//Inicializar repos
	userRepo := repositories.NewUserRepository(database.DB)

	//Inicializar handler
	userService := services.NewUserService(userRepo)

	//Inicializar servidor
	userHandler := handlers.NewUserHandler(userService)

	app := server.NewApp()

	app.Get("/health", health)
	app.Post("/signup", userHandler.SignUpHandler)

	err := app.RunServer(config.Port)

	if err != nil {
		log.Fatal("Error al iniciar el servirdor", err)
	}
}
