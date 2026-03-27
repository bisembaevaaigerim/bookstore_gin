package main

import (
	"bookstore_gin/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/{id}", handlers.GetBook)
	r.PUT("/books/{id}", handlers.UpdateBook)
	r.DELETE("/books/{id}", handlers.DeleteBook)

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	log.Println("Server running on http://localhost:8081")
	r.Run(":8081")
}
