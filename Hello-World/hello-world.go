package main

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name + "!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/:name", indexHandler)
	router.Run(":5000")
}
