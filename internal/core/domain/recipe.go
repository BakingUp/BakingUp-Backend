package domain

type Recipe struct {
	RecipeID   string `json:"recipe_id"`
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

type RecipeIngredient struct {
    IngredientName      string 		`json:"ingredient_name"`
    IngredientURL    	string 		`json:"ingredient_url"`
    IngredientQuantity 	string 		`json:"ingredient_quantity"`
    IngredientPrice  	float64    	`json:"ingredient_price"`
}

type RecipeDetail struct {
    Status            int          			`json:"status"`
    RecipeName        string       			`json:"recipe_name"`
    RecipeURL         []string       		`json:"recipe_url"`
    TotalTime         string       			`json:"total_time"`
    Servings          int          			`json:"servings"`
    Stars             int          			`json:"stars"`
    NumOfOrder        int          			`json:"num_of_order"`
    RecipeIngredients []RecipeIngredient 	`json:"recipe_ingredients"`
    InstructionURL    []string       		`json:"instruction_url"`
    InstructionSteps  []string     			`json:"instruction_steps"`
}