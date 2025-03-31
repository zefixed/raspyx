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
        "/api/v1/groups": {
            "get": {
                "description": "Get all groups from database",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Getting groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "response": {
                                            "$ref": "#/definitions/dto.GetGroupsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new group in the database and returns its uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Creating a new group",
                "parameters": [
                    {
                        "description": "Group number",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "response": {
                                            "$ref": "#/definitions/dto.CreateGroupResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/groups/number/{number}": {
            "get": {
                "description": "Get group from database with given number",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Getting group by number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "response": {
                                            "$ref": "#/definitions/models.Group"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/groups/uuid/{uuid}": {
            "get": {
                "description": "Get group from database with given uuid",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Getting group by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "response": {
                                            "$ref": "#/definitions/models.Group"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update group in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Updating group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Group",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/groups/{uuid}": {
            "delete": {
                "description": "Deleting existing group from the database",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Deleting existing group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateGroupRequest": {
            "type": "object",
            "required": [
                "group"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "example": "221-352"
                }
            }
        },
        "dto.CreateGroupResponse": {
            "type": "object",
            "properties": {
                "uuid": {
                    "type": "string",
                    "example": "c555b9e8-0d7a-11f0-adcd-20114d2008d9"
                }
            }
        },
        "dto.GetGroupsResponse": {
            "type": "object",
            "required": [
                "groups"
            ],
            "properties": {
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Group"
                    }
                }
            }
        },
        "dto.UpdateGroupRequest": {
            "type": "object",
            "required": [
                "group"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "example": "221-352"
                }
            }
        },
        "models.Group": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string",
                    "example": "221-352"
                },
                "uuid": {
                    "type": "string",
                    "example": "c555b9e8-0d7a-11f0-adcd-20114d2008d9"
                }
            }
        },
        "v1.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "Error"
                }
            }
        },
        "v1.ResponseOK": {
            "type": "object",
            "properties": {
                "response": {},
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Raspyx",
	Description:      "API for schedules",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
