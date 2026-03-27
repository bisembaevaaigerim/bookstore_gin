package handlers

import (
	"bookstore_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var authors = make(map[int]models.Author)
var authorNextID = 1

func GetAuthors(c *gin.Context) {
	authorsList := []models.Author{}
	for _, author := range authors {
		authorsList = append(authorsList, author)
	}
	c.JSON(http.StatusOK, authorsList)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(author.Name) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "name is required"})
		return
	}
	author.ID = authorNextID
	authorNextID++
	authors[author.ID] = author
	c.JSON(http.StatusOK, author)
}
