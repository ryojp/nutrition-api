package main

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

type FirstHit struct {
	Nutrition Nutrition `json:"nutrition"`
	SearchResult
}
