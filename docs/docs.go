// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

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
        "/api/v1/cube/audit": {
            "get": {
                "description": "query audit log from es",
                "tags": [
                    "audit"
                ],
                "summary": "query audit log",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "endTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "eventName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "resourceName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "responseStatus",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sourceIpAddress",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "startTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "userName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/audit.EsResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.ErrorInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/cube/audit/export": {
            "get": {
                "description": "query and export audit log from es",
                "tags": [
                    "audit"
                ],
                "summary": "export audit log",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "endTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "eventName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "resourceName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "responseStatus",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sourceIpAddress",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "startTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "userName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.ErrorInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/cube/healthz": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"hello\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "audit.EsResult": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Event"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "errcode.ErrorInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.Event": {
            "type": "object",
            "properties": {
                "apiAction": {
                    "type": "string"
                },
                "apiVersion": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "errorCode": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "eventName": {
                    "type": "string"
                },
                "eventTime": {
                    "type": "integer"
                },
                "eventType": {
                    "type": "string"
                },
                "eventVersion": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                },
                "requestMethod": {
                    "type": "string"
                },
                "requestParameters": {
                    "type": "string"
                },
                "resourceList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ResourceData"
                    }
                },
                "responseElements": {
                    "type": "string"
                },
                "responseStatus": {
                    "type": "integer"
                },
                "sourceIpAddress": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "userAgent": {
                    "type": "string"
                },
                "userIdentity": {
                    "$ref": "#/definitions/v1.UserIdentity"
                }
            }
        },
        "v1.ResourceData": {
            "type": "object",
            "properties": {
                "resourceId": {
                    "type": "string"
                },
                "resourceName": {
                    "type": "string"
                },
                "resourceType": {
                    "type": "string"
                }
            }
        },
        "v1.UserIdentity": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger KubeCube-Audit API",
	Description: "This is KubeCube-Audit api documentation.",
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
