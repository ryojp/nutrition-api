package main

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

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

// registerScrapeNutrition sets the nutrition scraper callback
func registerScrapeNutrition(collector *colly.Collector, nutrition *Nutrition) {
	collector.OnHTML("table#nutrition-label table tbody", scrapeNutrition(nutrition))
}

// registerScrapeSearchResults sets the search-results scraper callback
func registerScrapeSearchResults(collector *colly.Collector, searchResults *[]SearchResult) {
	collector.OnHTML("a.table_item_name", scrapeSearchResults(searchResults))
}

// scrapeSearchResults returns a handler that fills in the given searchResults array
func scrapeSearchResults(searchResults *[]SearchResult) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href == "" {
			return
		} else if href[0] == '/' { // then resolve the relative path
			href = baseUrl + href
		}
		name := e.Text
		*searchResults = append(*searchResults, SearchResult{
			Name: name,
			Url:  href,
		})
	}
}

// scrapeNutrition returns a handler that fills in the given nutrition struct
func scrapeNutrition(nutrition *Nutrition) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
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
	}
}
