# Nutrition API
This API returns nutrition data from [nutritionvalue.org](https://nutritionvalue.org)

It uses [Gin](https://github.com/gin-gonic/gin) for the API server and [Colly](https://github.com/gocolly/colly) for **caching** and **throttling**.


## Usage

### Docker (Pre-built image)

```sh
docker run --rm -it -p 8080:8080 -v $(pwd)/.cache/:/app/.cache/ ryojpn/nutrition-api
```

### Docker (Building on your own)
* Build an image with `docker build -t nutrition-api ./`
* Run it with `docker run --rm -it -p 8080:8080 -v $(pwd)/.cache/:/app/.cache/ nutrition-api`

### Go
Install [Go](https://go.dev/) and run `go run ./cmd/nutrition-api`


## Multi-platform building

1. Create a builder
```sh
docker buildx create --name armamd
```

2. Select the created builder
```sh
docker buildx use armamd
```

3. Build an image targeting `linux/arm64` and `linux/amd64`
```sh
docker buildx build --push --platform linux/amd64,linux/arm64 -t ryojpn/nutrition-api ./
```


## Routes
###  GET /nutrition?q={keyword}
This endpoint searches and visits the initial hit result.

`curl "http://localhost:8080/nutrition?q=bagel" | jq` (`jq` formats the output)
```json
{
  "nutrition": {
    "portion_size": "131 g",
    "calories": "346 kcal",
    "fat": "1.7 g",
    "saturated_fat": "0.5 g",
    "sodium": "553 mg",
    "carbs": "69 g",
    "fiber": "2.1 g",
    "sugar": "11 g",
    "protein": "14 g",
    "vitamin_d": "0 mcg",
    "calcium": "138 mg",
    "iron": "4.7 mg",
    "potassium": "140 mg"
  },
  "name": "Bagel",
  "url": "https://www.nutritionvalue.org/Bagel_51180010_nutritional_value.html"
}
```



###  GET /nutrition?url={url}
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

`curl "http://localhost:8080/nutrition?url=https://www.nutritionvalue.org/Banana%2C_raw_63107010_nutritional_value.html" | jq`
```json
{
  "portion_size": "225 g",
  "calories": "200 kcal",
  "fat": "0.7 g",
  "saturated_fat": "0.3 g",
  "sodium": "2.3 mg",
  "carbs": "51 g",
  "fiber": "5.9 g",
  "sugar": "28 g",
  "protein": "2.5 g",
  "vitamin_d": "0 mcg",
  "calcium": "11 mg",
  "iron": "0.6 mg",
  "potassium": "806 mg"
}
```


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

`curl "http://localhost:8080/search?q=salmon" | jq`
```json
[
  {
    "name": "Salmon, raw",
    "url": "https://www.nutritionvalue.org/Salmon%2C_raw_26137100_nutritional_value.html"
  },
  {
    "name": "Salmon loaf",
    "url": "https://www.nutritionvalue.org/Salmon_loaf_27250080_nutritional_value.html"
  },
  {
    "name": "Salmon salad",
    "url": "https://www.nutritionvalue.org/Salmon_salad_27450030_nutritional_value.html"
  },
  {
    "name": "Salmon, dried",
    "url": "https://www.nutritionvalue.org/Salmon%2C_dried_26137170_nutritional_value.html"
  },
  {
    "name": "Salmon, canned",
    "url": "https://www.nutritionvalue.org/Salmon%2C_canned_26137180_nutritional_value.html"
  },
  {
    "name": "Salmon, smoked",
    "url": "https://www.nutritionvalue.org/Salmon%2C_smoked_26137190_nutritional_value.html"
  },
  {
    "name": "Salmon cake or patty",
    "url": "https://www.nutritionvalue.org/Salmon_cake_or_patty_27250070_nutritional_value.html"
  },
  {
    "name": "Salmon cake sandwich",
    "url": "https://www.nutritionvalue.org/Salmon_cake_sandwich_27550120_nutritional_value.html"
  },
  {
    "name": "Salmon soup, cream style",
    "url": "https://www.nutritionvalue.org/Salmon_soup%2C_cream_style_28355350_nutritional_value.html"
  },
  {
    "name": "Salmon, steamed or poached",
    "url": "https://www.nutritionvalue.org/Salmon%2C_steamed_or_poached_26137160_nutritional_value.html"
  },
  {
    "name": "Fish, raw, chum, salmon",
    "url": "https://www.nutritionvalue.org/Fish%2C_raw%2C_chum%2C_salmon_nutritional_value.html"
  },
  {
    "name": "Fish, raw, pink, salmon",
    "url": "https://www.nutritionvalue.org/Fish%2C_raw%2C_pink%2C_salmon_nutritional_value.html"
  },
  {
    "name": "Salmon, no added fat, fried, coated",
    "url": "https://www.nutritionvalue.org/Salmon%2C_no_added_fat%2C_fried%2C_coated_26137143_nutritional_value.html"
  },
  {
    "name": "Fish, raw, chinook, salmon",
    "url": "https://www.nutritionvalue.org/Fish%2C_raw%2C_chinook%2C_salmon_nutritional_value.html"
  },
  {
    "name": "Fish, raw, sockeye, salmon",
    "url": "https://www.nutritionvalue.org/Fish%2C_raw%2C_sockeye%2C_salmon_nutritional_value.html"
  },
  {
    "name": "Lomi salmon",
    "url": "https://www.nutritionvalue.org/Lomi_salmon_27450310_nutritional_value.html"
  },
  {
    "name": "Salmon",
    "url": "https://www.nutritionvalue.org/Salmon_1028841_nutritional_value.html"
  },
  {
    "name": "Salmon, made with oil, fried, coated",
    "url": "https://www.nutritionvalue.org/Salmon%2C_made_with_oil%2C_fried%2C_coated_26137140_nutritional_value.html"
  },
  {
    "name": "Salmon, no added fat, baked or broiled",
    "url": "https://www.nutritionvalue.org/Salmon%2C_no_added_fat%2C_baked_or_broiled_26137123_nutritional_value.html"
  },
  {
    "name": "Fish, raw, wild, coho, salmon",
    "url": "https://www.nutritionvalue.org/Fish%2C_raw%2C_wild%2C_coho%2C_salmon_nutritional_value.html"
  }
]
```

