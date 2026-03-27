package handlers

import (
	"bookstore_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var categories = make(map[int]models.Category)
var categoryNextID = 1

func GetCategories(c *gin.Context) {
	list := []models.Category{}
	for _, cat := range categories {
		list = append(list, cat)
	}
	c.JSON(http.StatusOK, list)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(category.Name) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "name is required"})
		return
	}
	category.ID = categoryNextID
	categoryNextID++
	categories[category.ID] = category
	c.JSON(http.StatusCreated, category)
}
