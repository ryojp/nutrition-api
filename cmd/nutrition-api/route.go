package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// setupRouter registers all the routes
func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/search", handleSearch)
	r.GET("/view", handleView)

	return r
}

// handleSearch returns an array of nutrition pages
func handleSearch(c *gin.Context) {
	q := c.Query("q")
	visit := c.Query("visit")

	queryUrl := baseUrl + "/search.php?food_query=" + q

	var searchResults []SearchResult
	registerScrapeSearchResults(collector, &searchResults)

	collector.Visit(queryUrl)
	if visit == "" {
		c.JSON(http.StatusOK, searchResults)
	} else {
		n, err := strconv.Atoi(visit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'visit' must be an integer",
			})
		}
		n = (n + len(searchResults)) % len(searchResults)
		var nutrition Nutrition
		registerScrapeNutrition(collector, &nutrition)
		collector.Visit(searchResults[n].Url)
		c.JSON(http.StatusOK, nutrition)
	}
}

// handleView returns a nutrition response by visiting the given url
func handleView(c *gin.Context) {
	url := c.Query("url")
	if !strings.Contains(url, baseUrl) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid url",
		})
	}

	var nutrition Nutrition
	registerScrapeNutrition(collector, &nutrition)

	collector.Visit(url)
	c.JSON(http.StatusOK, nutrition)
}
