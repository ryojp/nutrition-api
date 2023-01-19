package main

import (
	"github.com/gocolly/colly/v2"
)

var baseUrl = "https://www.nutritionvalue.org"
var collector *colly.Collector

func main() {
	initCollector()
	r := setupRouter()
	r.Run(":8080")
}
