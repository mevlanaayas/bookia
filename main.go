package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mevlanaayas/bookia/controllers"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	_ = router.Run(":" + port) // listen and serve on 0.0.0.0:8080

	router.GET("/api/me/books", controllers.GetBooksFor)
	router.GET("/api/me/words", controllers.GetWordsFor)

}
