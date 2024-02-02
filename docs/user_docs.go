// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplateuser = `{
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
        "/user/activate": {
            "get": {
                "description": "激活账号",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "激活账号",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/bundle/get": {
            "get": {
                "description": "获取用户密钥包",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户密钥包",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/bundle/modify": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "修改用户密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户密钥包",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ModifyUserSecretBundleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/clients/get": {
            "get": {
                "description": "获取该用户当前登录的所有客户端",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取该用户当前登录的所有客户端",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "查询用户信息",
                "produces": [
                    "application/json"
                ],
                "summary": "查询用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status 用户状态 (0=未知状态, 1=正常状态, 2=被禁用, 3=已删除, 4=锁定状态) RelationStatus 用户关系状态 (0=不是好友, 1=是好友, 2=黑名单)",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/info/modify": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "修改用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/key/set": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "设置用户公钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "设置用户公钥",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetPublicKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "description": "退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "退出登录",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/password/modify": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "修改用户密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户密码",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/search": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "搜索用户",
                "produces": [
                    "application/json"
                ],
                "summary": "搜索用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status 用户状态 (0=未知状态, 1=正常状态, 2=被禁用, 3=已删除, 4=锁定状态) RelationStatus 用户关系状态 (0=不是好友, 1=是好友, 2=黑名单)",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/system/key/get": {
            "get": {
                "description": "获取系统pgp公钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取系统pgp公钥",
                "parameters": [
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "指定根据id还是邮箱类型查找",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.LogoutRequest": {
            "type": "object",
            "properties": {
                "login_number": {
                    "type": "integer"
                }
            }
        },
        "model.ModifyUserSecretBundleRequest": {
            "type": "object",
            "required": [
                "secret_bundle"
            ],
            "properties": {
                "secret_bundle": {
                    "type": "string"
                }
            }
        },
        "model.PasswordRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "old_password",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "public_key": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.SetPublicKeyRequest": {
            "type": "object",
            "required": [
                "public_key"
            ],
            "properties": {
                "public_key": {
                    "type": "string"
                }
            }
        },
        "model.UserInfoRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "model.UserInfoResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "login_number": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "relation_status": {
                    "$ref": "#/definitions/model.UserRelationStatus"
                },
                "signature": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.UserStatus"
                },
                "tel": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.UserRelationStatus": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-comments": {
                "UserRelationStatusBlacked": "黑名单 拉黑对方了",
                "UserRelationStatusFriend": "好友关系",
                "UserRelationStatusUnknown": "不是好友"
            },
            "x-enum-varnames": [
                "UserRelationStatusUnknown",
                "UserRelationStatusFriend",
                "UserRelationStatusBlacked"
            ]
        },
        "model.UserStatus": {
            "type": "integer",
            "enum": [
                0
            ],
            "x-enum-comments": {
                "UserStatusUnknown": "未知状态"
            },
            "x-enum-varnames": [
                "UserStatusUnknown"
            ]
        }
    }
}`

// SwaggerInfouser holds exported Swagger Info so clients can modify it
var SwaggerInfouser = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "coss-user服务",
	Description:      "",
	InfoInstanceName: "user",
	SwaggerTemplate:  docTemplateuser,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfouser.InstanceName(), SwaggerInfouser)
}
