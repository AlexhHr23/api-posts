package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexhHr23/gopost-api/server"
)

type Post struct {
	Id      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var posts []Post
var nextId = 1

func GetPosts(c *server.Context) {
	err := c.JSON(http.StatusOK, posts)
	if err != nil {
		http.Error(c.RWriter, "Error al codificar JSON", http.StatusInternalServerError)
	}
}

func CreatetPost(c *server.Context) {

	var post Post

	err := c.BindJson(post)

	if err != nil {
		http.Error(c.RWriter, "Error al decodificar JSON", http.StatusBadRequest)
	}

	post.Id = nextId
	nextId++
	posts = append(posts, post)

	err = c.JSON(http.StatusCreated, post)
	if err != nil {
		http.Error(c.RWriter, "Error al codificar JSON", http.StatusInternalServerError)
	}
}

func UpdatetPost(c *server.Context) {

	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	var updatedPost Post
	err := c.BindJson(updatedPost)
	if err != nil {
		http.Error(c.RWriter, "Error al decodificar JSON", http.StatusBadRequest)
	}

	for i := range posts {
		if posts[i].Id == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			c.JSON(http.StatusOK, posts[i])
			return
		}
	}

	http.Error(c.RWriter, "Post no encontrado", http.StatusNotFound)
}

func DeletePost(c *server.Context) {
	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for i := range posts {
		if posts[i].Id == id {
			posts = append(posts[:i], posts[i+1:]...)
			http.Error(c.RWriter, "Post borrado correctamente", http.StatusOK)
			return
		}
	}

	http.Error(c.RWriter, "Post no encontrado", http.StatusNotFound)
}

func GetPostById(c *server.Context) {
	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.Id == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}

	http.Error(c.RWriter, "Post no encontrado", http.StatusNotFound)
}
