package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Fiction"},
	{ID: 2, Name: "Mystery"},
	{ID: 3, Name: "Science Fiction"},
}

var nextCategoryID = 4

func GetCategories(c *gin.Context) {
	nameFilter := c.Query("name")
	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	var filteredCategories []models.Category
	for _, category := range categories {
		if nameFilter == "" || strings.Contains(strings.ToLower(category.Name), strings.ToLower(nameFilter)) {
			filteredCategories = append(filteredCategories, category)
		}
	}

	total := len(filteredCategories)
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{"categories": []models.Category{}, "total": 0})
		return
	}

	start := (page - 1) * limit
	if start >= total {
		c.JSON(http.StatusOK, gin.H{"categories": []models.Category{}, "total": total})
		return
	}

	end := start + limit
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{"categories": filteredCategories[start:end], "total": total})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	category.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, category)

	c.JSON(http.StatusCreated, category)
}
