// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
      "description": "Musahit Harita API",
      "title": "Musahit Harita API",
      "contact": {
        "name": "Acik Kaynak"
      },
      "license": {
        "name": "Apache 2.0",
        "url": "https://github.com/acikkaynak/musahit-harita-backend/blob/main/LICENSE"
      },
      "version": "1.0"
    },
    "host": "localhost:80",
    "paths": {
      "/feeds": {
        "get": {
          "produces": ["application/json"],
          "tags": ["Feed"],
          "summary": "Get Feeds",
          "responses": {
            "200": {
              "description": "OK",
              "schema": {
                "$ref": "#/definitions/feeds.Response"
              }
            }
          }
        }
      },
      "/feeds/mock": {
        "get": {
          "produces": ["application/json"],
          "tags": ["Feed"],
          "summary": "Get Feeds mock",
          "responses": {
            "200": {
              "description": "OK",
              "schema": {
                "$ref": "#/definitions/feeds.Response"
              }
            }
          }
        }
      },
      "/feed/mock/{neighborhoodId}": {
        "get": {
          "produces": ["application/json"],
          "tags": ["Feed"],
          "summary": "Get Feed Detail mock",
          "parameters": [
            {
              "type": "integer",
              "description": "neighborhoodId",
              "name": "neighborhoodId",
              "in": "path",
              "required": true
            }
          ],
          "responses": {
            "200": {
              "description": "OK",
              "schema": {
                "$ref": "#/definitions/feeds.FeedDetailResponse"
              }
            }
          }
        }
      },
      "/feed/{neighborhoodId}": {
        "get": {
          "produces": ["application/json"],
          "tags": ["Feed"],
          "summary": "Get Feed Detail",
          "parameters": [
            {
              "type": "integer",
              "description": "neighborhoodId",
              "name": "neighborhoodId",
              "in": "path",
              "required": true
            }
          ],
          "responses": {
            "200": {
              "description": "OK",
              "schema": {
                "$ref": "#/definitions/feeds.FeedDetailResponse"
              }
            }
          }
        }
      },
      "/volunteer-form": {
        "post": {
          "produces": ["application/json"],
          "tags": ["Volunteer"],
          "summary": "Volunteer Form",
          "parameters": [
            {
              "type": "integer",
              "description": "neighborhood identifier",
              "name": "neighborhoodId",
              "required": true
            },
            {
              "type": "string",
              "description": "Your name",
              "name": "name",
              "required": true
            },
            {
              "type": "string",
              "description": "Your Surname",
              "name": "surname",
              "required": true
            },
            {
              "type": "string",
              "description": "Your email address",
              "name": "email",
              "required": true
            },
            {
              "type": "string",
              "description": "Your phone number",
              "name": "phone",
              "required": true
            },
            {
              "type": "boolean",
              "description": "KVKK approval",
              "name": "kvkk",
              "required": true
            }
          ],
          "responses": {
            "200": {
              "description": "OK",
              "schema": {
                "$ref": "#/definitions/volunteer.Ok"
              }
            },
            "400": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error.ErrorResponse"
              }
            }
          }
        }
      }
    },
    "definitions": {
      "feeds.Feed": {
        "type": "object",
        "properties": {
          "neighborhood_id": {
            "type": "integer"
          },
          "volunteer_data": {
            "type": "integer"
          }
        }
      },
      "feeds.Response": {
        "type": "object",
        "properties": {
          "count": {
            "type": "integer"
          },
          "results": {
            "type": "array",
            "items": {
              "$ref": "#/definitions/feeds.Feed"
            }
          }
        }
      },
      "feeds.FeedDetailResponse": {
        "type": "object",
        "properties": {
          "neighborhoodId": {
            "type": "integer"
          },
          "lastUpdateTime": {
            "type": "string"
          },
          "intensity": {
            "type": "integer"
          },
          "details": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "volunteer.Ok": {
        "type": "object",
        "properties": {
          "ok": {
            "type": "boolean"
          }
        }
      },
      "error.ErrorResponse": {
        "type": "array",
        "items": {
          "type": "object",
          "properties": {
            "field": {
              "type": "string"
            },
            "message": {
              "type": "string"
            },
            "value": {
              "type": "string"
            }
          }
        }
      }
    }
  }
  `

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Musahit Harita API",
	Description:      "Musahit Harita API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
