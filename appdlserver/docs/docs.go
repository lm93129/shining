// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-28 10:42:29.2594541 +0800 CST m=+0.054852501

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
        "contact": {
            "name": "lm93129",
            "url": "https://fs.tn",
            "email": "lm93129@qq.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/appFile/appInfo/{id}": {
            "get": {
                "description": "获取单个安装包详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "获取单个安装包详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "安装包ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appFile/appInfos": {
            "get": {
                "description": "获取安装包列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "获取安装包列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页大小",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "所属项目,不填则查询所有项目",
                        "name": "project_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "安装包类型,不填则查询项目下所有安装包",
                        "name": "version_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appFile/deleteFile": {
            "delete": {
                "description": "删除安装包",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "删除安装包",
                "parameters": [
                    {
                        "description": "安装包id数组",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appfile.DeleteApp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appFile/plist/{id}": {
            "get": {
                "description": "IOS获取plist文件",
                "produces": [
                    "text/xml"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "IOS获取plist文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "安装包ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appfileparser.sparseBundleHeader"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appFile/project": {
            "get": {
                "description": "APP安装包统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "APP安装包统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "star_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "结束事件",
                        "name": "end_time",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appFile/uploadFile": {
            "post": {
                "description": "上传安装包文件",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "APP安装包相关"
                ],
                "summary": "上传安装包文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "安装包所属项目ID",
                        "name": "ProjectId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "安装包更新说明",
                        "name": "UpdateDescription",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "安装包文件",
                        "name": "upload",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "err_code：50001 数据库操作失败 err_code: 40001 参数错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "健康检查接口，如果返回200则表示成功",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "健康检查"
                ],
                "summary": "健康检查",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appfile.DeleteApp": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "description": "要删除的安装包ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "appfileparser.Assets": {
            "type": "object",
            "properties": {
                "kind": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "appfileparser.Item": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/appfileparser.Assets"
                    }
                },
                "metadata": {
                    "type": "object",
                    "$ref": "#/definitions/appfileparser.Metadata"
                }
            }
        },
        "appfileparser.Metadata": {
            "type": "object",
            "properties": {
                "bundleIdentifier": {
                    "type": "string"
                },
                "bundleVersion": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "appfileparser.sparseBundleHeader": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/appfileparser.Item"
                    }
                }
            }
        },
        "serializer.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "localhost:3000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "shining安装包管理系统",
	Description: "shining一个APP安装包管理系统",
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
