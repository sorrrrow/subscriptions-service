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
        "/subscriptions": {
            "get": {
                "description": "Возвращает список всех подписок из БД",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение всех подписок",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Subscription"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новую подписку в БД",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создание подписки",
                "parameters": [
                    {
                        "description": "Данные подписки",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Subscription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Subscription"
                        }
                    },
                    "400": {
                        "description": "Неверный JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/subscriptions/summary": {
            "get": {
                "description": "Возвращает сумму цен подписок по user_id, дате и имени сервиса",
                "produces": [
                    "application/json"
                ],
                "summary": "Сумма подписок по фильтрам",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Название сервиса",
                        "name": "service_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата начала периода (YYYY-MM-DD)",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Дата конца периода (YYYY-MM-DD)",
                        "name": "to",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка параметров",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/subscriptions/{id}": {
            "put": {
                "description": "Обновляет существующую подписку по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Обновление подписки",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID подписки",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Обновлённые данные подписки",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Subscription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Subscription"
                        }
                    },
                    "400": {
                        "description": "Неверный JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Подписка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет подписку по ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Удаление подписки",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID подписки",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно удалено",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Подписка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Subscription": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Subscriptions API",
	Description:      "API для управления онлайн-подписками пользователей",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
