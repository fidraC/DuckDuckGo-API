package main

import (
	"os"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"github.com/acheong08/DuckDuckGo-API/duckduckgo"
	"github.com/acheong08/DuckDuckGo-API/types"
)

func main() {
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
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
			ctx.JSON(400, gin.H{"error": err.Error(), "details": "Could not bind JSON"})
			return
		}
		// Ensure query is set
		if search.Query == "" {
			ctx.JSON(400, gin.H{"error": "Query is required"})
			return
		}
		// Get results
		results, err := duckduckgo.Get_results(search)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Limit
		if search.Limit > 0 && search.Limit < len(results) {
			results = results[:search.Limit]
		}
		// Return results
		ctx.JSON(200, results)
	})
	handler.GET("/search", func(ctx *gin.Context) {
		// Map request to Search struct
		var search types.Search
		// Get query
		search.Query = ctx.Query("query")
		// Get region
		search.Region = ctx.Query("region")
		// Get time range
		search.TimeRange = ctx.Query("time_range")
		if search.Query == "" {
			ctx.JSON(400, gin.H{"error": "Query is required"})
			return
		}
		// Get limit and check if it's a number
		limit := ctx.Query("limit")
		if limit != "" {
			if _, err := strconv.Atoi(limit); err != nil {
				ctx.JSON(400, gin.H{"error": "Limit must be a number"})
				return
			}
			search.Limit, _ = strconv.Atoi(limit)
		}
		// Get results
		results, err := duckduckgo.Get_results(search)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Shorten results to limit if limit is set
		if search.Limit > 0 && search.Limit < len(results) {
			results = results[:search.Limit]
		}
		// Return results
		ctx.JSON(200, results)
	})

	endless.ListenAndServe(HOST+":"+PORT, handler)
}
