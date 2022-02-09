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

func FindBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
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

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	config.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
