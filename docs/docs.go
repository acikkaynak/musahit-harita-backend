// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Acik Kaynak"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/acikkaynak/musahit-harita-backend/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/feed/{neighborhoodId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
                "summary": "Get Feed Detail",
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
        "/feeds": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
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
        "/mock/feed/{neighborhoodId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
                "summary": "Get Feed Detail mock",
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
        "/mock/feeds": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
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
        "feeds.FeedDetail": {
            "type": "object",
            "properties": {
                "ballotBoxNos": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "buildingName": {
                    "type": "string"
                }
            }
        },
        "feeds.FeedDetailResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/feeds.FeedDetail"
                    }
                },
                "intensity": {
                    "type": "integer"
                },
                "lastUpdateTime": {
                    "type": "string"
                },
                "neighborhoodId": {
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
        }
    }
}`

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
