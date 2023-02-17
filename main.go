package main

import (
	"github.com/akhilesh-ingle-ge/pkg/controller"
	"github.com/akhilesh-ingle-ge/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDB()
	router := gin.Default()

	router.GET("/books", controller.GetAllBooks)
	router.GET("/books/:id", controller.GetBookById)
	router.POST("/books", controller.CreateBooks)
	router.PATCH("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)

	// log.Fatal(http.ListenAndServe(":8080", router))
	router.Run(":8080")
}
