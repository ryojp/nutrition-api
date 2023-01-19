# Nutrition API
This API returns nutrition data from `nutritionvalue.org`

It uses [Gin](https://github.com/gin-gonic/gin) for the API server and [Colly](https://github.com/gocolly/colly) for caching and throttling.

## Usage
Install [Go](https://go.dev/) and run `go run ./cmd/nutrition-api`

## Routes
###  GET /search?q={keyword}
The response type is
```typescript
[
    {
        name: string,
        url: string,
    }
]
```

`curl "http://localhost:8080/search?q=salmon"`
```json
[{"name":"Salmon, raw","url":"https://www.nutritionvalue.org/Salmon%2C_raw_26137100_nutritional_value.html"},{"name":"Salmon loaf","url":"https://www.nutritionvalue.org/Salmon_loaf_27250080_nutritional_value.html"},{"name":"Salmon salad","url":"https://www.nutritionvalue.org/Salmon_salad_27450030_nutritional_value.html"}]
```

###  GET /search?q={keyword}&visit={index}
With `visit={index:int}` query parameter, the nutrition data of the `index`-th search result is returned, internally using `GET /view` API.
`curl "http://localhost:8080/search?q=onion&visit=0"`
```json
{"portion_size":"160 g","calories":"64 kcal","fat":"0.2 g","saturated_fat":"0.1 g","sodium":"6.4 mg","carbs":"15 g","fiber":"2.7 g","sugar":"6.8 g","protein":"1.8 g","vitamin_d":"0 mcg","calcium":"37 mg","iron":"0.3 mg","potassium":"234 mg"}
```

###  GET /view?url={url}
In response, the unit and the number are separated with a whitespace.
```typescript
[
    {
        portion_size: string,
        calories: string,
        fat: string,
        saturated_fat: string,
        sodium: string,
        carbs: string,
        fiber: string,
        sugar: string,
        protein: string,
        vitamin_d: string,
        calcium: string,
        iron: string,
        potassium: string,
    }
]
```

`curl "http://localhost:8080/view?url=https://www.nutritionvalue.org/Banana%2C_raw_63107010_nutritional_value.html"`
```json
{"portion_size":"225 g","calories":"200 kcal","fat":"0.7 g","saturated_fat":"0.3 g","sodium":"2.3 mg","carbs":"51 g","fiber":"5.9 g","sugar":"28 g","protein":"2.5 g","vitamin_d":"0 mcg","calcium":"11 mg","iron":"0.6 mg","potassium":"806 mg"}
```


