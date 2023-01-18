package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

type Nutrition struct {
	PortionSize  string `json:"portion_size"`
	Calories     string `json:"calories"`
	Fat          string `json:"fat"`
	SaturatedFat string `json:"saturated_fat"`
	Sodium       string `json:"sodium"`
	Carbs        string `json:"carbs"`
	Fiber        string `json:"fiber"`
	Sugar        string `json:"sugar"`
	Protein      string `json:"protein"`
	VitaminD     string `json:"vitamin_d"`
	Calcium      string `json:"calcium"`
	Iron         string `json:"iron"`
	Potassium    string `json:"potassium"`
}

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

// handleView returns a nutrition response by visiting the given url
func handleView(c *gin.Context) {
	url := c.Query("url")
	if !strings.Contains(url, baseUrl) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid url",
		})
	}

	var nutrition Nutrition

	collector.OnHTML("table#nutrition-label table tbody", func(e *colly.HTMLElement) {
		nutrition.PortionSize = e.ChildText("#serving-size")
		nutrition.Calories = e.ChildText("#calories") + " kcal"

		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(strings.ReplaceAll(s.Text(), "\u00a0", " "))
			line := strings.Split(text, "\n")[0]
			matches := regexp.MustCompile(`(([0-9]*[.])?[0-9]+)\s?(\w+)`).FindStringSubmatch(line)
			if len(matches) < 2 {
				return
			}
			num, unit := matches[1], matches[len(matches)-1]
			val := num + " " + unit
			switch {
			case strings.Contains(line, "Total Fat"):
				nutrition.Fat = val
			case strings.Contains(line, "Saturated Fat"):
				nutrition.SaturatedFat = val
			case strings.Contains(line, "Sodium"):
				nutrition.Sodium = val
			case strings.Contains(line, "Total Carbohydrate"):
				nutrition.Carbs = val
			case strings.Contains(line, "Dietary Fiber"):
				nutrition.Fiber = val
			case strings.Contains(line, "Sugar"):
				nutrition.Sugar = val
			case strings.Contains(line, "Protein"):
				nutrition.Protein = val
			case strings.Contains(line, "Vitamin D"):
				nutrition.VitaminD = val
			case strings.Contains(line, "Calcium"):
				nutrition.Calcium = val
			case strings.Contains(line, "Iron"):
				nutrition.Iron = val
			case strings.Contains(line, "Potassium"):
				nutrition.Potassium = val
			}
		})
	})

	collector.Visit(url)
	c.JSON(http.StatusOK, nutrition)
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
