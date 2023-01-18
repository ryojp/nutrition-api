package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

var baseUrl = "https://www.nutritionvalue.org"
var collector *colly.Collector

type SearchResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/search", handleSearch)

	return r
}

func handleSearch(c *gin.Context) {
	q := c.Query("q")

	queryUrl := baseUrl + "/search.php?food_query=" + q

	var searchResults []SearchResult

	collector.OnHTML("a.table_item_name", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href == "" {
			return
		} else if href[0] == '/' { // then resolve the relative path
			href = baseUrl + href
		}
		name := e.Text
		searchResults = append(searchResults, SearchResult{
			Name: name,
			Url:  href,
		})
	})

	collector.Visit(queryUrl)
	c.JSON(http.StatusOK, searchResults)
}

func initCollector() {
	collector = colly.NewCollector(
		colly.CacheDir("./.cache/nutritionvalue"),
		colly.AllowURLRevisit(), // disable the already-visited check
	)
	collector.Limit(&colly.LimitRule{
		RandomDelay: 1 * time.Second,
	})
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)
}

func main() {
	initCollector()
	r := setupRouter()
	r.Run(":8080")
}
