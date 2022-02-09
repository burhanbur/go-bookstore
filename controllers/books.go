package controllers

import (
	"net/http"

	config "bookstore/config"
	models "bookstore/models"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": books,
		},
	)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	config.DB.Create(&book)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": book,
		},
	)
}
