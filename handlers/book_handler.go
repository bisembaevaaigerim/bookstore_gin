package handlers

import (
	"bookstore_gin/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var books = make(map[int]models.Book)
var bookNextID = 1

func GetBooks(c *gin.Context) {
	categoryFilter := strings.ToLower(c.Query("category"))
	authorIDFilter, _ := strconv.Atoi(c.Query("author_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	list := []models.Book{}
	for _, b := range books {
		if authorIDFilter > 0 && b.AuthorID != authorIDFilter {
			continue
		}
		if categoryFilter != "" {
			cat, ok := categories[b.CategoryID]
			if !ok || strings.ToLower(cat.Name) != categoryFilter {
				continue
			}
		}
		list = append(list, b)
	}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].ID > list[j].ID {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	total := len(list)
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"books": list[start:end],
	})
}
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(book.Title) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "title is required"})
		return
	}
	if _, ok := authors[book.AuthorID]; !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "author not found"})
		return
	}
	if _, ok := categories[book.CategoryID]; !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "category not found"})
		return
	}
	if book.Price <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "price must be greater than 0"})
		return
	}
	book.ID = bookNextID
	bookNextID++
	books[book.ID] = book
	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	book, ok := books[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	if _, ok := books[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(input.Title) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "title is required"})
		return
	}
	if _, ok := authors[input.AuthorID]; !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "author not found"})
		return
	}
	if _, ok := categories[input.CategoryID]; !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "category not found"})
		return
	}
	if input.Price <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "price must be greater than 0"})
		return
	}
	input.ID = id
	books[id] = input
	c.JSON(http.StatusOK, input)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	if _, ok := books[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}
