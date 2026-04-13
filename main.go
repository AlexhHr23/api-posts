package main

import (
	"context"
	// "fmt"
	"log"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/database"

	// "github.com/AlexhHr23/gopost-api/models"
	"github.com/AlexhHr23/gopost-api/repositories"
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

	postRepo := repositories.NewPostRepository(database.DB)

	//LIstar posts
	posts, err := postRepo.FindAll(context.Background())

	if err != nil {
		log.Fatal("Error al obtener posts: ", err)
	}

	for _, post := range posts {
		log.Printf("Post ID: %d, UserID: %d, Title: %s\n", post.ID, post.UserID, post.Title)
	}

	//Listar post por id
	post, err := postRepo.FindById(context.Background(), 2)

	if err != nil {
		log.Fatal("Error al obtener post por ID: ", err)
	}

	log.Printf("Post obtenido por ID: %d, Title: %s,\n", post.ID, post.Title)

	//Post de user ID
	userPosts, err := postRepo.FindByUserId(context.Background(), 2)

	if err != nil {
		log.Fatal("Error al obtener posts por user ID: ", err)
	}

	for _, post := range userPosts {
		log.Printf("Post de user ID 2 - Post ID: %d , Title: %s\n", post.ID, post.Title)
	}

	//Actualizar
	// post.Title = "Titulo actualizado"
	// post.Content = "Contenido actualizado"
	// err = postRepo.Update(context.Background(), post)

	// if err != nil {
	// 	log.Fatal("Error al actualizar el post: ", err)
	// }

	// log.Printf("Post actualizado: ID: %d, Title: %s\n", post.ID, post.Title)

	//Eliminar
	err = postRepo.Delete(context.Background(), 3)

	if err != nil {
		log.Fatal("Error al eliminar el post: ", err)
	} else {
		log.Println("Post eliminado con exito")
	}

	// app.Get("/health", health)

	// app.Post("/posts", handlers.CreatetPost)
	// app.Get("/posts", handlers.GetPosts)
	// app.Put("/posts", handlers.UpdatetPost)
	// app.Get("/posts/{id}", handlers.GetPostById)
	// app.Delete("/posts/{id}", handlers.DeletePost)

	err = app.RunServer(config.Port)

	if err != nil {
		log.Fatal("Error al iniciar el servirdor", err)
	}
}
