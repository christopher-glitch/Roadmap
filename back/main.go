package main

import (

	//"github.com/joho/godotenv"
  //_ "github.com/go-sql-driver/mysql"

	"roadmap/database"

	"github.com/gin-gonic/gin"
	
)

func main() {
	result := database.Connect()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": result,
		})
	})

	router.Run(":8080")
}
