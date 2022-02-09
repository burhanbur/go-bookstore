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
	router.POST("/books", controllers.CreateBook)

	fmt.Println("CONNECTED TO PORT 8080")
	log.Fatal(router.Run(":8080"))
}
