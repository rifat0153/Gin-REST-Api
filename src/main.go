package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Home Page",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}


func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func main() {
	fmt.Print("Hello")

	r := gin.Default()

	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", PathParameters)


	r.Run()
}