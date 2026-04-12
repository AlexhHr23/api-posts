package main

import (
	"fmt"
	"net/http"

	"github.com/AlexhHr23/gopost-api/handlers"
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

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func main() {

	mux := http.NewServeMux()
	// mux.HandleFunc("GET /hola/{name}/{age}", hola)
	// mux.HandleFunc("GET /hola/info", hola)

	mux.HandleFunc("GET /health", health)

	mux.HandleFunc("GET /posts", handlers.GetPosts)
	mux.HandleFunc("POST /posts", handlers.CreatetPost)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostById)
	mux.HandleFunc("PUT /posts/{id}", handlers.UpdatetPost)
	mux.HandleFunc("DELETE /posts/{id}", handlers.DeletePost)

	fmt.Println("Servirdor iniciendo en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
