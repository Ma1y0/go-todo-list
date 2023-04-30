package controllers

import (
	"net/http"

	"github.com/Ma1y0/go-todo-list/models"
	"github.com/gin-gonic/gin"
)

// Get all authors
// GET /authors
func GetAllAuthors(c *gin.Context) {
    var authors []models.Author 
    models.DB.Find(&authors)
    c.JSON(http.StatusOK, gin.H{"data": authors})
}

// Get author by id
// GET //authors/:id
func GetAuthorById(c *gin.Context) {
    var author models.Author

    if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": author})

}

// Create author
// POST /authors
func CreateAuthor(c *gin.Context) {
    var input models.Author
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    author := models.Author{
       Name: input.Name, 
    }

    result := models.DB.Create(&author)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"dara": author})
}

// Update a author
type updateAuthorInput struct {
    Name string `json:"name"`
    Bio string `json:"bio"`
    Photo string `json:"photo"`
}
// PATCH /authors/:id
func UpdateAuthor(c *gin.Context) {
    var author models.Author

    if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        return
    }

    var input updateAuthorInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&author).Updates(input)
    
    c.JSON(http.StatusOK, gin.H{"data": author})
}

// Delete a author
// DELETE /authors/:id
func DeleteAuthor(c *gin.Context) {
    var author models.Author
    if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        return
    }

    models.DB.Delete(&author)
}
