package main

import (
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"github.com/acheong08/DuckDuckGo-API/internal/types"
)

func main() {
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	if HOST == "" {
		HOST = "127.0.0.1"
	}
	if PORT == "" {
		PORT = "8080"
	}
	handler := gin.Default()
	handler.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	handler.POST("/search", func(ctx *gin.Context) {
		// Map request to Search struct
		var search types.Search
		if err := ctx.ShouldBindJSON(&search); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	})

	endless.ListenAndServe(HOST+":"+PORT, handler)
}
