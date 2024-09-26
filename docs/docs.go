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
        }
    },
    "definitions": {
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
