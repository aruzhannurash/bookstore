package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books []models.Book
var nextBookID = 1

func GetBooks(c *gin.Context) {
	categoryFilter := c.Query("category")
	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 5
	}

	var filteredBooks []models.Book
	for _, book := range books {
		if categoryFilter == "" || strconv.Itoa(book.CategoryID) == categoryFilter {
			filteredBooks = append(filteredBooks, book)
		}
	}

	total := len(filteredBooks)
	totalPages := (total + limit - 1) / limit

	if total == 0 || (page-1)*limit >= total {
		c.JSON(http.StatusOK, gin.H{"books": []models.Book{}, "total": total, "totalPages": totalPages})
		return
	}

	start := (page - 1) * limit
	end := start + limit
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{"books": filteredBooks[start:end], "total": total, "totalPages": totalPages})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book.ID = nextBookID
	nextBookID++
	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			var updatedBook models.Book
			if err := c.ShouldBindJSON(&updatedBook); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			updatedBook.ID = id
			books[i] = updatedBook

			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
