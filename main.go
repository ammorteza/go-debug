package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/rest/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Morteza Amzajerdi",
		})
	})
	server.Run(":9999")
}
