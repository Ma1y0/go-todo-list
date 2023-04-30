package main

import (
	"net/http"

	"github.com/Ma1y0/go-todo-list/models"
    "github.com/Ma1y0/go-todo-list/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectToDatabase()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
    })
    // Books router
    router.GET("/books", controllers.GetBooks)
    router.GET("/books/:id", controllers.GetAuthorById)
    router.POST("/books", controllers.CreateBook)
    // Authors router
    router.GET("/authors", controllers.GetAllAuthors)
    router.GET("/authors/:id", controllers.GetAuthorById)
    router.POST("/authors", controllers.CreateAuthor)
    router.PATCH("/authors/:id", controllers.UpdateAuthor)
    router.DELETE("/authors/:id", controllers.DeleteAuthor)
 

    router.Run()
}
