package main

import (
	"fmt"
	"log"

	config "bookstore/config"
	controllers "bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.InitDb()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	fmt.Println("CONNECTED TO PORT 8080")
	log.Fatal(router.Run(":8080"))
}
