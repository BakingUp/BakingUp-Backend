// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ingredient/deleteIngredient": {
            "delete": {
                "description": "Delete an ingredient by using ingredient id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingredient"
                ],
                "summary": "Delete an ingredient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient ID",
                        "name": "ingredient_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot delete an ingredient",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/ingredient/deleteIngredientBatchNote": {
            "delete": {
                "description": "Delete ingredient batch note by ingredient note ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingredient"
                ],
                "summary": "Delete ingredient batch note",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient Note ID",
                        "name": "ingredient_note_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot delete ingredient batch note",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/ingredient/getAllIngredients": {
            "get": {
                "description": "Get all ingredients by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingredient"
                ],
                "summary": "Get all ingredients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.IngredientList"
                        }
                    },
                    "400": {
                        "description": "Cannot get all ingredients",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/ingredient/getIngredientDetail": {
            "get": {
                "description": "Get ingredient details by ingredient ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingredient"
                ],
                "summary": "Get ingredient details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient ID",
                        "name": "ingredient_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.IngredientDetail"
                        }
                    },
                    "400": {
                        "description": "Cannot get ingredient detail",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/ingredient/getIngredientStockDetail": {
            "get": {
                "description": "Get ingredient stock details by ingredient stock ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingredient"
                ],
                "summary": "Get ingredient stock details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient stock ID",
                        "name": "ingredient_stock_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.IngredientStockDetail"
                        }
                    },
                    "400": {
                        "description": "Cannot get ingredient stock detail",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/recipe/deleteRecipe": {
            "delete": {
                "description": "Delete a recipe by using recipe id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Delete a recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot delete a recipe",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/recipe/getRecipeDetail": {
            "get": {
                "description": "Get recipe details by recipe ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Get recipe details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.RecipeDetail"
                        }
                    },
                    "400": {
                        "description": "Cannot get recipe detail",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/changeColorExpired": {
            "put": {
                "description": "Change the color of expired icon by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Change the color of expired icon",
                "parameters": [
                    {
                        "description": "Change Color Expired Icon",
                        "name": "change_color_expired_icon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChangeExpirationDateSetting"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully change the color of expiration icon",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot change the color of expiration icon",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/changeFixCost": {
            "put": {
                "description": "Change the fix cost by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Change the fix cost",
                "parameters": [
                    {
                        "description": "Change Fix Cost",
                        "name": "change_fix_cost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChangeFixCostSetting"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully change the fix cost",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot change the fix cost",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/changeLanguage": {
            "put": {
                "description": "Change the application language by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Change the application language",
                "parameters": [
                    {
                        "description": "Change Language",
                        "name": "change_language",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChangeUserLanguage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully change the language.",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot change the language",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/deleteAccount": {
            "delete": {
                "description": "Delete an account by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Delete an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully delete an account",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot delete an account",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/getColorExpired": {
            "get": {
                "description": "Get the color of expired icon by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Get the color of expired icon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.ExpirationDateSetting"
                        }
                    },
                    "400": {
                        "description": "Cannot get the color of expiration icon",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/getFixCost": {
            "get": {
                "description": "Get the fix cost by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Get the fix cost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.FixCostSetting"
                        }
                    },
                    "400": {
                        "description": "Cannot get the fix cost",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/settings/getLanguage": {
            "get": {
                "description": "Get the application language by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "settings"
                ],
                "summary": "Get the application language",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.UserLanguage"
                        }
                    },
                    "400": {
                        "description": "Cannot get the language",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/stock/deleteStock": {
            "delete": {
                "description": "Delete a stock by recipe id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock"
                ],
                "summary": "Delete a stock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot delete a stock",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/stock/getAllStocks": {
            "get": {
                "description": "Get all stocks by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock"
                ],
                "summary": "Get all stocks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.StockList"
                        }
                    },
                    "400": {
                        "description": "Cannot get all stocks",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/stock/getStockDetail": {
            "get": {
                "description": "Get stock details by recipe ID and user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock"
                ],
                "summary": "Get stock details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/domain.StockDetail"
                        }
                    },
                    "400": {
                        "description": "Cannot get stock detail",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/user/editUserInfo": {
            "put": {
                "description": "Edit user information by using user information request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Edit user information",
                "parameters": [
                    {
                        "description": "Edit User Info",
                        "name": "edit_user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ManageUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully edit the user information.",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Cannot edit the user information",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ChangeExpirationDateSetting": {
            "type": "object",
            "properties": {
                "black_expiration_date": {
                    "type": "integer"
                },
                "red_expiration_date": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "yellow_expiration_date": {
                    "type": "integer"
                }
            }
        },
        "domain.ChangeFixCostSetting": {
            "type": "object",
            "properties": {
                "advertising": {
                    "type": "number"
                },
                "electricity": {
                    "type": "number"
                },
                "gas": {
                    "type": "number"
                },
                "insurance": {
                    "type": "number"
                },
                "note": {
                    "type": "string"
                },
                "other": {
                    "type": "number"
                },
                "rent": {
                    "type": "number"
                },
                "salaries": {
                    "type": "number"
                },
                "subscriptions": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                },
                "water": {
                    "type": "number"
                }
            }
        },
        "domain.ChangeUserLanguage": {
            "type": "object",
            "properties": {
                "language": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.ExpirationDateSetting": {
            "type": "object",
            "properties": {
                "black_expiration_date": {
                    "type": "integer"
                },
                "red_expiration_date": {
                    "type": "integer"
                },
                "yellow_expiration_date": {
                    "type": "integer"
                }
            }
        },
        "domain.FixCostSetting": {
            "type": "object",
            "properties": {
                "advertising": {
                    "type": "number"
                },
                "electricity": {
                    "type": "number"
                },
                "gas": {
                    "type": "number"
                },
                "insurance": {
                    "type": "number"
                },
                "note": {
                    "type": "string"
                },
                "other": {
                    "type": "number"
                },
                "rent": {
                    "type": "number"
                },
                "salaries": {
                    "type": "number"
                },
                "subscriptions": {
                    "type": "number"
                },
                "water": {
                    "type": "number"
                }
            }
        },
        "domain.Ingredient": {
            "type": "object",
            "properties": {
                "expiration_status": {
                    "type": "string"
                },
                "ingredient_id": {
                    "type": "string"
                },
                "ingredient_name": {
                    "type": "string"
                },
                "ingredient_url": {
                    "type": "string"
                },
                "quantity": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "domain.IngredientDetail": {
            "type": "object",
            "properties": {
                "ingredient_less_than": {
                    "type": "integer"
                },
                "ingredient_name": {
                    "type": "string"
                },
                "ingredient_quantity": {
                    "type": "string"
                },
                "ingredient_url": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock_amount": {
                    "type": "integer"
                },
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Stock"
                    }
                }
            }
        },
        "domain.IngredientList": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Ingredient"
                    }
                }
            }
        },
        "domain.IngredientNote": {
            "type": "object",
            "properties": {
                "ingredient_note": {
                    "type": "string"
                },
                "ingredient_note_id": {
                    "type": "string"
                },
                "note_created_at": {
                    "type": "string"
                }
            }
        },
        "domain.IngredientStockDetail": {
            "type": "object",
            "properties": {
                "day_before_expire": {
                    "type": "string"
                },
                "ingredient_brand": {
                    "type": "string"
                },
                "ingredient_eng_name": {
                    "type": "string"
                },
                "ingredient_price": {
                    "type": "string"
                },
                "ingredient_quantity": {
                    "type": "string"
                },
                "ingredient_stock_url": {
                    "type": "string"
                },
                "ingredient_supplier": {
                    "type": "string"
                },
                "ingredient_thai_name": {
                    "type": "string"
                },
                "notes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.IngredientNote"
                    }
                }
            }
        },
        "domain.ManageUserRequest": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "store_name": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.RecipeDetail": {
            "type": "object",
            "properties": {
                "instruction_steps": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "instruction_url": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "num_of_order": {
                    "type": "integer"
                },
                "recipe_ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.RecipeIngredient"
                    }
                },
                "recipe_name": {
                    "type": "string"
                },
                "recipe_url": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "servings": {
                    "type": "integer"
                },
                "stars": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "total_time": {
                    "type": "string"
                }
            }
        },
        "domain.RecipeIngredient": {
            "type": "object",
            "properties": {
                "ingredient_name": {
                    "type": "string"
                },
                "ingredient_price": {
                    "type": "number"
                },
                "ingredient_quantity": {
                    "type": "string"
                },
                "ingredient_url": {
                    "type": "string"
                }
            }
        },
        "domain.Stock": {
            "type": "object",
            "properties": {
                "expiration_date": {
                    "type": "string"
                },
                "expiration_status": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "quantity": {
                    "type": "string"
                },
                "stock_id": {
                    "type": "string"
                },
                "stock_url": {
                    "type": "string"
                }
            }
        },
        "domain.StockDetail": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "lst_status": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "sell_by_date": {
                    "type": "string"
                }
            }
        },
        "domain.UserLanguage": {
            "type": "object",
            "properties": {
                "language": {
                    "type" : "string"
                }
            }
        }
        "domain.StockItem": {
            "type": "object",
            "properties": {
                "lst": {
                    "type": "integer"
                },
                "lst_status": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "selling_price": {
                    "type": "number"
                },
                "stock_id": {
                    "type": "string"
                },
                "stock_name": {
                    "type": "string"
                },
                "stock_url": {
                    "type": "string"
                }
            }
        },
        "domain.StockList": {
            "type": "object",
            "properties": {
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.StockItem"
                    }
                }
            }
        },
        "http.response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "BakingUp Backend API",
	Description:      "This is the BakingUp Backend API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
