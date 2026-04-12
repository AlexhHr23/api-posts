package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Post struct {
	Id      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var posts []Post
var nextId = 1

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(posts)

	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
	}
}

func CreatetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
	}

	post.Id = nextId
	nextId++
	posts = append(posts, post)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(post)
	fmt.Println("err", err)
	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
	}
}

func UpdatetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	var updatedPost Post
	err := json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
	}

	for i := range posts {
		if posts[i].Id == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Post no encontrado", http.StatusNotFound)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for i := range posts {
		if posts[i].Id == id {
			posts = append(posts[:i], posts[i+1:]...)
			http.Error(w, "Post borrado correctamente", http.StatusOK)
			return
		}
	}

	http.Error(w, "Post no encontrado", http.StatusNotFound)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.Id == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, "Post no encontrado", http.StatusNotFound)
}
