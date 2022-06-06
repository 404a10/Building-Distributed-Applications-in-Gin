package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:firstName, attr`
	LastName  string   `xml:lastName, attr`
}

func indexHandler(c *gin.Context) {
	c.XML(200, Person{FirstName: "Debra",
		LastName: "T Patterson"})
}

func main() {
	router := gin.Default()
	router.GET("/", indexHandler)
	router.Run(":5000")
}
