package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/husnizuhdi/test-app-engine/cloudstorage"
)

func main() {
	name := os.Getenv("NAME")
	hello := "Hello " + name + " and the World!"

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": hello,
		})
	})

	r.POST("/upload", cloudstorage.HandleFileUploadToBucket)

	r.Run()
}
