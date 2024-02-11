package main

import (

	//"github.com/joho/godotenv"
  //_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Air!",
		})
	})

	router.Run(":8080")
}
