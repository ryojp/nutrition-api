package main

import (
	"net/http"
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
	r.GET("/nutrition", handleNutrition)

	return r
}

// handleSearch returns an array of nutrition pages
func handleSearch(c *gin.Context) {
	q := c.Query("q")

	queryUrl := baseUrl + "/search.php?food_query=" + q

	var searchResults []SearchResult
	registerScrapeSearchResults(collector, &searchResults)

	collector.Visit(queryUrl)
	c.JSON(http.StatusOK, searchResults)
}

// handleNutrition dispatches a handler and returns the nutrition
func handleNutrition(c *gin.Context) {
	q := c.Query("q")
	url := c.Query("url")

	if url != "" {
		handleView(c)
	} else if q != "" {
		handleFirstHit(c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect query parameter",
		})
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

// handleFirstHit returns the nutrition of the first hit result
func handleFirstHit(c *gin.Context) {
	q := c.Query("q")
	queryUrl := baseUrl + "/search.php?food_query=" + q

	var searchResults []SearchResult
	registerScrapeSearchResults(collector, &searchResults)
	collector.Visit(queryUrl)

	if len(searchResults) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no result found from the given search query",
		})
		return
	}

	var nutrition Nutrition
	registerScrapeNutrition(collector, &nutrition)
	collector.Visit(searchResults[0].Url)
	c.JSON(http.StatusOK, SearchResultWithNutrition{
		SearchResult: searchResults[0],
		Nutrition:    nutrition,
	})
}
