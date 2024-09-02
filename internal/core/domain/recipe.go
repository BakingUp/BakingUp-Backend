package domain

type Recipe struct {
	RecipeName string `json:"recipe_name"`
	RecipeImg  string `json:"recipe_img"`
	TotalTime  string `json:"total_time"`
	Servings   int    `json:"servings"`
	Stars      int    `json:"stars"`
	NumOfOrder int    `json:"num_of_order"`
}

type RecipeList struct {
	Recipes []Recipe `json:"recipes"`
}
