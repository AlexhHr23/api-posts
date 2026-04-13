package main

import (
	"log"

	"github.com/AlexhHr23/gopost-api/config"
	"github.com/AlexhHr23/gopost-api/handlers"
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

	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /hola/{name}/{age}", hola)
	// mux.HandleFunc("GET /hola/info", hola)

	// mux.HandleFunc("GET /health", health)

	// mux.HandleFunc("GET /posts", handlers.GetPosts)
	// mux.HandleFunc("POST /posts", handlers.CreatetPost)
	// mux.HandleFunc("GET /posts/{id}", handlers.GetPostById)
	// mux.HandleFunc("PUT /posts/{id}", handlers.UpdatetPost)
	// mux.HandleFunc("DELETE /posts/{id}", handlers.DeletePost)

	// fmt.Println("Servirdor iniciendo en http://localhost:8080")
	// http.ListenAndServe(":8080", mux)

	config := config.LoadConfig()

	app := server.NewApp()

	app.Get("/health", health)

	app.Post("/posts", handlers.CreatetPost)
	app.Get("/posts", handlers.GetPosts)
	app.Put("/posts", handlers.UpdatetPost)
	app.Get("/posts/{id}", handlers.GetPostById)
	app.Delete("/posts/{id}", handlers.DeletePost)

	err := app.RunServer(config.Port)

	if err != nil {
		log.Fatal("Error al iniciar el servirdor", err)
	}
}
