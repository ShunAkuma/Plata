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
        "/quotes": {
            "patch": {
                "description": "Updating the quote in the background",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quotes"
                ],
                "summary": "Update Quotes Rate",
                "parameters": [
                    {
                        "description": "Сurrency code",
                        "name": "CurrencyCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Response"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/quotes/:id": {
            "get": {
                "description": "Get quotes rate by id from redis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quotes"
                ],
                "summary": "Get quotes rate by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "quotes",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Response"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/quotes/currency/:currency": {
            "get": {
                "description": "Get last quotes with time and rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quotes"
                ],
                "summary": "Get last quotes rate with time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Currency code",
                        "name": "CurrencyCode",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Response"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Response": {
            "description": "Response model from quotes server",
            "type": "object",
            "properties": {
                "message": {
                    "description": "Message from server\nexample: \"Something went wrong\"",
                    "type": "string"
                },
                "resultObj": {
                    "description": "Generalized Response Object from Server\nResponse model that can store an error or response"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Plata backend",
	Description:      "Quotation server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
