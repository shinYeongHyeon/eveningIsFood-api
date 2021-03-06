// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "EveningIsFood",
            "email": "eveningisfood@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/foodCard": {
            "post": {
                "description": "음식점카드를 생성하는 API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foodShopCard"
                ],
                "summary": "음식점카드 생성",
                "parameters": [
                    {
                        "description": "음식점카드 생성 Request Body",
                        "name": "foodCard",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/foodCardControllerCommandDto.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "음식점카드 생성 Response",
                        "schema": {
                            "$ref": "#/definitions/foodCardControllerCommandDto.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/foodCardControllerCommandDto.CreateResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/foodCardControllerCommandDto.CreateResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "foodCardControllerCommandDto.CreateRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "음식점 주소",
                    "type": "string",
                    "example": "서울특별시 용산구 남영동 80-4"
                },
                "name": {
                    "description": "음식점 이름",
                    "type": "string",
                    "example": "초원"
                }
            }
        },
        "foodCardControllerCommandDto.CreateResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "SUCCESS"
                },
                "foodCard": {
                    "$ref": "#/definitions/foodCardControllerCommandDto.FoodCard"
                }
            }
        },
        "foodCardControllerCommandDto.CreateResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "foodCardControllerCommandDto.FoodCard": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "음식점 주소",
                    "type": "string"
                },
                "name": {
                    "description": "음식점 이름",
                    "type": "string"
                },
                "uuid": {
                    "description": "음식점 고유 ID",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0",
	Host:             "eveningisfood.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "EveningIsFood API",
	Description:      "예로부터 약속은 저녁에 이루어 졌고, 역사는 술 속에서 만들어졌습니다.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
