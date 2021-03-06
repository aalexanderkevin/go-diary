// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/diary": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create Specific Diary by Date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diary"
                ],
                "summary": "Diary services",
                "parameters": [
                    {
                        "description": "Create Diary",
                        "name": "diary",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Diary"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/diary/{year}/{quarter}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get Content of the Diary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diary"
                ],
                "summary": "Diary services",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quarter",
                        "name": "quarter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Diary"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates user and provides a JWT to Authorize API calls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "login services",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JWT"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Logout user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logout"
                ],
                "summary": "Logout services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "Register services",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Diary": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "This is content of the diary"
                },
                "date": {
                    "type": "string",
                    "example": "2021-01-12"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string",
                    "example": "invalid JSON format"
                }
            }
        },
        "models.JWT": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTA4OTY5MjYsInVzZXJuYW1lIjoibGFsYSIsInV1aWQiOiI5YTA3YTIwYi01MTYwLTQ4N2ItYTBlYS1iMzBkZjM3NmMyMjcifQ.3kMJSHB-pMSjRovvVPU1O1p6Y04qgLDaJKr1ONPtkvY"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd!23"
                },
                "username": {
                    "type": "string",
                    "example": "lala"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string",
                    "example": "1996-04-14"
                },
                "email": {
                    "type": "string",
                    "example": "lala@gmailcom"
                },
                "name": {
                    "type": "string",
                    "example": "lala lili"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd!23"
                },
                "username": {
                    "type": "string",
                    "example": "lala"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
