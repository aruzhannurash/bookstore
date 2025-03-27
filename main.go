 package main

import (
	"bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	books := r.Group("/books")
	{
		books.GET("/", handlers.GetBooks)
		books.POST("/", handlers.CreateBook)
		books.GET("/:id", handlers.GetBookByID)
		books.PUT("/:id", handlers.UpdateBook)
		books.DELETE("/:id", handlers.DeleteBook)
	}

	authors := r.Group("/authors")
	{
		authors.GET("/", handlers.GetAuthors)
		authors.POST("/", handlers.CreateAuthor)
	}

	categories := r.Group("/categories")
	{
		categories.GET("/", handlers.GetCategories)
		categories.POST("/", handlers.CreateCategory)
	}

	r.Run(":8080")
}
