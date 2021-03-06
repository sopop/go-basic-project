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
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "首页",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/html"
                ],
                "summary": "首页",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/goods_del": {
            "post": {
                "description": "删除商品信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除商品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apitoken验证",
                        "name": "ApiToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/goods_detail": {
            "post": {
                "description": "获取商品信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apitoken验证",
                        "name": "ApiToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/service.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Goods"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/goods_lists": {
            "get": {
                "description": "展示商品",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/html"
                ],
                "summary": "商品页",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "获取商品列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apitoken验证",
                        "name": "ApiToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "每页条数",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/service.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Goods"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/goods_operate": {
            "post": {
                "description": "添加、编辑商品信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加、编辑商品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apitoken验证",
                        "name": "ApiToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品描述",
                        "name": "detail",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Goods": {
            "type": "object",
            "properties": {
                "cat_id": {
                    "description": "分类ID",
                    "type": "integer"
                },
                "create_time": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "create_time_str": {
                    "description": "创建时间格式化",
                    "type": "string"
                },
                "detail": {
                    "description": "商品描述",
                    "type": "string"
                },
                "id": {
                    "description": "商品ID",
                    "type": "integer"
                },
                "name": {
                    "description": "商品名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "update_time": {
                    "description": "更新时间",
                    "type": "integer"
                }
            }
        },
        "service.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "为0时表示成功",
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "service.ResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
