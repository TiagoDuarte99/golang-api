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
        "/users": {
            "post": {
                "description": "This endpoint allows you to create a new user with the provided data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SignupRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflit",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SignupRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "name",
                "password",
                "user_type"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "minLength": 6,
                    "example": "SenhaForte123!"
                },
                "email": {
                    "type": "string",
                    "example": "tiago@example.com"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2,
                    "example": "Tiago Duarte"
                },
                "password": {
                    "type": "string",
                    "minLength": 6,
                    "example": "SenhaForte123!"
                },
                "user_type": {
                    "type": "string",
                    "example": "USER"
                }
            }
        },
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
