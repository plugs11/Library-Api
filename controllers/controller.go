package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"manan.tola/models"
	"manan.tola/utils"
	"manan.tola/verification"
)

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	utils.ParseBody(c.Request, CreateBook)
	if verification.CheckduplicateId(CreateBook.Id) {
		c.JSON(409, gin.H{"message": "existing Id"})
		return
	}
	if CreateBook.Title == "" || CreateBook.Author == "" || CreateBook.PublishedYear == "" || CreateBook.Genre == "" || CreateBook.Id == "" {
		c.JSON(409, gin.H{"error": "empty fields"})
		return
	} else {
		details := CreateBook.CreateBook()
		response, _ := json.Marshal(details)
		c.Writer.Write(response)
	}
} // connects to models and creates a new book

func GetAllBook(c *gin.Context) {
	books := models.GetAllBook()
	response, _ := json.Marshal(books)
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)

} // connects to models and provides data of all book

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	if !verification.CheckduplicateId(id) {
		c.JSON(409, gin.H{"message": "invalid id"})
		return
	}
	if id == "" {
		c.JSON(204, gin.H{"message": "empty fields"})
		return
	} else {
		bookDetails := models.GetBookByID(id)
		response, _ := json.Marshal(bookDetails)
		c.Data(http.StatusOK, "application/json; charset=utf-8", response)
	}

}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if !verification.CheckduplicateId(id) {
		c.JSON(409, gin.H{"message": "invalid id"})
		return
	}
	if id == "" {
		c.JSON(204, gin.H{"message": "empty fields"})
		return
	} else {
		var updatedBookDetails []models.Book
		UpdateBook := &models.Book{}
		utils.ParseBody(c.Request, UpdateBook)
		updatedBookDetails = models.UpdateBook(id, *UpdateBook, c)
		response, _ := json.Marshal(updatedBookDetails)
		if updatedBookDetails != nil {
			c.Data(http.StatusOK, "application/json; charset=utf-8", response)
		}
	}

}

func DeleteBook(c *gin.Context) { /////////////////DONE
	id := c.Param("id")
	if !verification.CheckduplicateId(id) {
		c.JSON(409, gin.H{"message": "invalid id"})
		return
	}
	if id == "" {
		c.JSON(204, gin.H{"message": "empty fields"})
		return
	} else {
		books := models.DeleteBook(id)
		response, _ := json.Marshal(books)
		c.Data(http.StatusOK, "application/json; charset=utf-8", response)
	}
} // connects to models and deletes a book
