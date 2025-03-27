package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var authors = []models.Author{
	{ID: 1, Name: "author 1"},
	{ID: 2, Name: "author 2"},
	{ID: 3, Name: "author 3"},
}

var nextAuthorID = 4

func GetAuthors(c *gin.Context) {
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

	var filteredAuthors []models.Author
	for _, author := range authors {
		if nameFilter == "" || strings.Contains(strings.ToLower(author.Name), strings.ToLower(nameFilter)) {
			filteredAuthors = append(filteredAuthors, author)
		}
	}

	total := len(filteredAuthors)
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{"authors": []models.Author{}, "total": 0})
		return
	}

	start := (page - 1) * limit
	if start >= total {
		c.JSON(http.StatusOK, gin.H{"authors": []models.Author{}, "total": total})
		return
	}

	end := start + limit
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{"authors": filteredAuthors[start:end]})
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	author.ID = nextAuthorID
	nextAuthorID++
	authors = append(authors, author)

	c.JSON(http.StatusCreated, author)
}
