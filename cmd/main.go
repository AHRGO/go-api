package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	
	_baseUrl := "/api"

	server.GET(_baseUrl + "/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8025")
}