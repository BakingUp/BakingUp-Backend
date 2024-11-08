package domain

import "time"

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
	IngredientName     string  `json:"ingredient_name"`
	IngredientURL      string  `json:"ingredient_url"`
	IngredientQuantity string  `json:"ingredient_quantity"`
	IngredientPrice    float64 `json:"ingredient_price"`
}

type RecipeDetail struct {
	Status            int                `json:"status"`
	RecipeName        string             `json:"recipe_name"`
	RecipeURL         []string           `json:"recipe_url"`
	TotalTime         string             `json:"total_time"`
	Servings          int                `json:"servings"`
	Stars             int                `json:"stars"`
	NumOfOrder        int                `json:"num_of_order"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients"`
	InstructionURL    []string           `json:"instruction_url"`
	InstructionSteps  string             `json:"instruction_steps"`
	HiddenCost        float64            `json:"hidden_cost"`
	LaborCost         float64            `json:"labor_cost"`
	ProfitMargin      float64            `json:"profit_margin"`
}

type AddRecipeIngredientRequest struct {
	IngredientID       string `json:"ingredient_id"`
	IngredientQuantity string `json:"ingredient_quantity"`
}

type AddRecipeRequest struct {
	UserID          string                       `json:"user_id"`
	RecipeEngName   string                       `json:"recipe_eng_name"`
	RecipeThaiName  string                       `json:"recipe_thai_name"`
	RecipeImg       []string                     `json:"recipe_img"`
	TotalHours      string                       `json:"total_hours"`
	TotalMins       string                       `json:"total_mins"`
	Servings        string                       `json:"servings"`
	Ingredients     []AddRecipeIngredientRequest `json:"ingredients"`
	InstructionImg  []string                     `json:"instruction_img"`
	EngInstruction  string                       `json:"eng_instruction"`
	ThaiInstruction string                       `json:"thai_instruction"`
}

type AddRecipePayload struct {
	UserID          string    `json:"user_id"`
	RecipeID        string    `json:"recipe_id"`
	RecipeEngName   string    `json:"recipe_eng_name"`
	RecipeThaiName  string    `json:"recipe_thai_name"`
	TotalTime       time.Time `json:"total_time"`
	Servings        int       `json:"servings"`
	EngInstruction  string    `json:"eng_instruction"`
	ThaiInstruction string    `json:"thai_instruction"`
}

type AddRecipeImagePayload struct {
	RecipeID   string `json:"ingredient_id"`
	ImgUrl     string `json:"img"`
	ImageIndex int    `json:"image_index"`
}

type AddRecipeInstructionImagePayload struct {
	RecipeID   string `json:"recipe_id"`
	ImgUrl     string `json:"img"`
	ImageIndex int    `json:"image_index"`
}

type AddRecipeIngredientPayload struct {
	RecipeID                 string  `json:"recipe_id"`
	IngredientID             string  `json:"ingredient_id"`
	RecipeIngredientQuantity float64 `json:"recipe_ingredient_quantity"`
}

type UpdateHiddenCostRequest struct {
	RecipeID   string `json:"recipe_id"`
	HiddenCost string `json:"hidden_cost"`
}

type UpdateLaborCostRequest struct {
	RecipeID  string `json:"recipe_id"`
	LaborCost string `json:"labor_cost"`
}

type UpdateProfitMarginRequest struct {
	RecipeID     string `json:"recipe_id"`
	ProfitMargin string `json:"profit_margin"`
}

type UpdateHiddenCostPayload struct {
	RecipeID   string  `json:"recipe_id"`
	HiddenCost float64 `json:"hidden_cost"`
}

type UpdateLaborCostPayload struct {
	RecipeID  string  `json:"recipe_id"`
	LaborCost float64 `json:"labor_cost"`
}

type UpdateProfitMarginPayload struct {
	RecipeID     string  `json:"recipe_id"`
	ProfitMargin float64 `json:"profit_margin"`
}

type EditRecipeRequest struct {
	RecipeID         string                          `json:"recipe_id"`
	RecipeEngName    string                          `json:"recipe_eng_name"`
	RecipeThaiName   string                          `json:"recipe_thai_name"`
	TotalHours       string                          `json:"total_hours"`
	TotalMins        string                          `json:"total_mins"`
	Servings         string                          `json:"servings"`
	EngInstruction   string                          `json:"eng_instruction"`
	ThaiInstruction  string                          `json:"thai_instruction"`
	Ingredients      []AddRecipeIngredientRequest    `json:"ingredients"`
}

type EditRecipePayload struct {
	RecipeID         string    `json:"recipe_id"`
	RecipeEngName    string    `json:"recipe_eng_name"`
	RecipeThaiName   string    `json:"recipe_thai_name"`
	TotalTime        time.Time `json:"total_time"`
	Servings         int       `json:"servings"`
	EngInstruction   string    `json:"eng_instruction"`
	ThaiInstruction  string    `json:"thai_instruction"`
}

type GetEditRecipeDetail struct {
	RecipeEngName   string                          `json:"recipe_eng_name"`
	RecipeThaiName  string                          `json:"recipe_thai_name"`
	TotalHours      string                          `json:"total_hours"`
	TotalMins       string                          `json:"total_mins"`
	Servings        string                          `json:"servings"`
	EngInstruction  string                          `json:"eng_instruction"`
	ThaiInstruction string                          `json:"thai_instruction"`
	Ingredients     []GetEditRecipeIngredientDetail `json:"ingredients"`
}

type GetEditRecipeIngredientDetail struct {
	IngredientID       string `json:"ingredient_id"`
	IngredientName     string `json:"ingredient_name"`
	IngredientURL      string `json:"ingredient_url"`
	IngredientQuantity string `json:"ingredient_quantity"`
	IngredientUnit     string `json:"ingredient_unit"`
}
