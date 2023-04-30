package controllers

import (
	"net/http"

	"github.com/Ma1y0/go-todo-list/models"
	"github.com/gin-gonic/gin"
)

// Gets all books
// GET /books
func GetBooks(c *gin.Context) {
    var books []models.Book
    models.DB.Find(&books)

    c.JSON(http.StatusOK, gin.H{"data": books})
}

// Gets book a by Id
func GetBookById(c *gin.Context) {
    var book models.Book

    if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create a book 
// POST /books
func CreateBook(c *gin.Context) {
    var input struct {
        Name        string         `json:"name" binding:"required"`
        AuthorID    uint           `json:"author_id" binding:"required"`
        Description string         `json:"description"`
        Cover       string         `json:"cover"`
    }
    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    var author models.Author
    result := models.DB.First(&author, input.AuthorID)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
        return
    }

    book := models.Book{
        Name: input.Name,
        Author: author,
    }

    result = models.DB.Create(&book)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": book})
}

// Update a book
type updateBookInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}
// PATCH /books/:id
func UpdateBook(c *gin.Context) {
    var input updateAuthorInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    var book models.Book

    if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
    }

    models.DB.Model(&book).Updates(&input)
}
