package main

var baseUrl = "https://www.nutritionvalue.org"

func main() {
	initCollector()
	r := setupRouter()
	r.Run(":8080")
}
