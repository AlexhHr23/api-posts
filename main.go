package main

import (

	// "fmt"
	"log"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/database"

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

	app := server.NewApp()

	// app.Get("/health", health)

	// app.Post("/posts", handlers.CreatetPost)
	// app.Get("/posts", handlers.GetPosts)
	// app.Put("/posts", handlers.UpdatetPost)
	// app.Get("/posts/{id}", handlers.GetPostById)
	// app.Delete("/posts/{id}", handlers.DeletePost)

	err := app.RunServer(config.Port)

	if err != nil {
		log.Fatal("Error al iniciar el servirdor", err)
	}
}
