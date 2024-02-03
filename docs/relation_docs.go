// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplaterelation = `{
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
        "/relation/delete_blacklist": {
            "post": {
                "description": "删除黑名单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "删除黑名单",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeleteBlacklistRequest"
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
        "/relation/dialog/show": {
            "post": {
                "description": "关闭或打开对话(action: 0:关闭对话, 1:打开对话)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dialog"
                ],
                "summary": "关闭或打开对话(action: 0:关闭对话, 1:打开对话)",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CloseOrOpenDialogRequest"
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
        "/relation/dialog/top": {
            "post": {
                "description": "是否置顶对话(action: 0:关闭取消置顶对话, 1:置顶对话)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dialog"
                ],
                "summary": "是否置顶对话(action: 0:关闭取消置顶对话, 1:置顶对话)",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TopOrCancelTopDialogRequest"
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
        "/relation/group_relation/admin/manage/join": {
            "post": {
                "description": "管理员管理加入群聊 action (0=拒绝, 1=同意)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "管理员管理加入群聊",
                "parameters": [
                    {
                        "description": "Action (0: rejected, 1: joined)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ManageJoinGroupRequest"
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
        "/relation/group_relation/admin/manage/remove": {
            "post": {
                "description": "将用户从群聊移除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "将用户从群聊移除",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RemoveUserFromGroupRequest"
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
        "/relation/group_relation/invite": {
            "post": {
                "description": "邀请加入群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "邀请加入群聊",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InviteGroupRequest"
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
        "/relation/group_relation/join": {
            "post": {
                "description": "加入群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "加入群聊",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.JoinGroupRequest"
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
        "/relation/group_relation/list": {
            "get": {
                "description": "群聊列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "群聊列表",
                "responses": {
                    "200": {
                        "description": "status 0:正常状态；1:被封禁状态；2:被删除状态",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/usersorter.CustomGroupData"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/relation/group_relation/manage_join": {
            "post": {
                "description": "用户管理加入群聊 action (0=拒绝, 1=同意)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "用户管理加入群聊",
                "parameters": [
                    {
                        "description": "Action (0: rejected, 1: joined)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ManageJoinGroupRequest"
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
        "/relation/group_relation/member": {
            "get": {
                "description": "群聊成员列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "群聊成员列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "群聊ID",
                        "name": "group_id",
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
        "/relation/group_relation/quit": {
            "post": {
                "description": "退出群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "退出群聊",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.QuitGroupRequest"
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
        "/relation/group_relation/request_list": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取用户的群聊申请列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "获取群聊申请列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status (0=申请中, 1=待通过, 2=已加入, 3=已删除, 4=被拒绝, 5=被封禁)",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.GroupRequestListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/relation/group_relation/silent": {
            "post": {
                "description": "设置群聊静默通知",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GroupRelation"
                ],
                "summary": "设置群聊静默通知",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetGroupSilentNotificationRequest"
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
        "/relation/user_relation/add_blacklist": {
            "post": {
                "description": "添加黑名单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "添加黑名单",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddBlacklistRequest"
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
        "/relation/user_relation/add_friend": {
            "post": {
                "description": "发送好友请求",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "发送好友请求",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendFriendRequest"
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
        "/relation/user_relation/blacklist": {
            "get": {
                "description": "黑名单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "黑名单",
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
        "/relation/user_relation/delete_friend": {
            "post": {
                "description": "删除好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "删除好友",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeleteFriendRequest"
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
        "/relation/user_relation/friend_list": {
            "get": {
                "description": "好友列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "好友列表",
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
        "/relation/user_relation/manage_friend": {
            "post": {
                "description": "管理好友请求  action (0=拒绝, 1=同意)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "管理好友请求",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ManageFriendRequest"
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
        "/relation/user_relation/request_list": {
            "get": {
                "description": "好友申请列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "好友申请列表",
                "responses": {
                    "200": {
                        "description": "UserStatus 申请状态 (0=申请中, 1=待通过, 2=已添加, 3=被拒绝, 4=已删除, 5=已拒绝)",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.UserRequestListResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/relation/user_relation/silent": {
            "post": {
                "description": "设置私聊静默通知",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "设置私聊静默通知",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetUserSilentNotificationRequest"
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
        "/relation/user_relation/switch/e2e/key": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "交换用户端到端公钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRelation"
                ],
                "summary": "交换用户端到端公钥",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwitchUserE2EPublicKeyRequest"
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
        }
    },
    "definitions": {
        "model.ActionEnum": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-comments": {
                "ActionAccepted": "同意",
                "ActionRejected": "拒绝"
            },
            "x-enum-varnames": [
                "ActionRejected",
                "ActionAccepted"
            ]
        },
        "model.AddBlacklistRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.CloseOrOpenDialogAction": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "CloseDialog",
                "OpenDialog"
            ]
        },
        "model.CloseOrOpenDialogRequest": {
            "type": "object",
            "required": [
                "dialog_id"
            ],
            "properties": {
                "action": {
                    "$ref": "#/definitions/model.CloseOrOpenDialogAction"
                },
                "dialog_id": {
                    "type": "integer"
                }
            }
        },
        "model.DeleteBlacklistRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.DeleteFriendRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.GroupRequestListResponse": {
            "type": "object",
            "properties": {
                "creator_id": {
                    "type": "string"
                },
                "group_avatar": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "group_name": {
                    "type": "string"
                },
                "group_status": {
                    "type": "integer"
                },
                "group_type": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "max_members_limit": {
                    "type": "integer"
                },
                "receiver_info": {
                    "$ref": "#/definitions/model.UserInfo"
                },
                "remark": {
                    "type": "string"
                },
                "sender_info": {
                    "$ref": "#/definitions/model.UserInfo"
                },
                "status": {
                    "$ref": "#/definitions/model.GroupRequestStatus"
                }
            }
        },
        "model.GroupRequestStatus": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-comments": {
                "Accepted": "已通过",
                "InvitationReceived": "邀请接收者",
                "InviteSender": "邀请发送者",
                "Pending": "等待",
                "Rejected": "已拒绝"
            },
            "x-enum-varnames": [
                "Pending",
                "Accepted",
                "Rejected",
                "InviteSender",
                "InvitationReceived"
            ]
        },
        "model.InviteGroupRequest": {
            "type": "object",
            "required": [
                "group_id",
                "member"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                },
                "member": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.JoinGroupRequest": {
            "type": "object",
            "required": [
                "group_id"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                }
            }
        },
        "model.ManageFriendRequest": {
            "type": "object",
            "required": [
                "request_id"
            ],
            "properties": {
                "action": {
                    "$ref": "#/definitions/model.ActionEnum"
                },
                "e2e_public_key": {
                    "type": "string"
                },
                "request_id": {
                    "type": "integer"
                }
            }
        },
        "model.ManageJoinGroupRequest": {
            "type": "object",
            "required": [
                "group_id",
                "id"
            ],
            "properties": {
                "action": {
                    "$ref": "#/definitions/model.ActionEnum"
                },
                "group_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.QuitGroupRequest": {
            "type": "object",
            "required": [
                "group_id"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                }
            }
        },
        "model.RemoveUserFromGroupRequest": {
            "type": "object",
            "required": [
                "group_id",
                "member"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                },
                "member": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
        "model.SendFriendRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "e2e_public_key": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.SetGroupSilentNotificationRequest": {
            "type": "object",
            "required": [
                "group_id"
            ],
            "properties": {
                "group_id": {
                    "description": "群ID",
                    "type": "integer"
                },
                "is_silent": {
                    "description": "是否开启静默通知",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.SilentNotificationType"
                        }
                    ]
                }
            }
        },
        "model.SetUserSilentNotificationRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "is_silent": {
                    "description": "是否开启静默通知",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.SilentNotificationType"
                        }
                    ]
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "string"
                }
            }
        },
        "model.SilentNotificationType": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-comments": {
                "IsSilent": "开启静默通知",
                "NotSilent": "静默通知关闭"
            },
            "x-enum-varnames": [
                "NotSilent",
                "IsSilent"
            ]
        },
        "model.SwitchUserE2EPublicKeyRequest": {
            "type": "object",
            "required": [
                "public_key",
                "user_id"
            ],
            "properties": {
                "public_key": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.TopOrCancelTopAction": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "CancelTopDialog",
                "TopDialog"
            ]
        },
        "model.TopOrCancelTopDialogRequest": {
            "type": "object",
            "required": [
                "dialog_id"
            ],
            "properties": {
                "action": {
                    "$ref": "#/definitions/model.TopOrCancelTopAction"
                },
                "dialog_id": {
                    "type": "integer"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "user_avatar": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.UserRequestListResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "receiver_id": {
                    "type": "string"
                },
                "receiver_info": {
                    "$ref": "#/definitions/model.UserInfo"
                },
                "remark": {
                    "type": "string"
                },
                "request_at": {
                    "type": "integer"
                },
                "sender_id": {
                    "type": "string"
                },
                "sender_info": {
                    "$ref": "#/definitions/model.UserInfo"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "usersorter.CustomGroupData": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "dialog_id": {
                    "type": "integer"
                },
                "group_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInforelation holds exported Swagger Info so clients can modify it
var SwaggerInforelation = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "coss-relation-bff服务",
	Description:      "",
	InfoInstanceName: "relation",
	SwaggerTemplate:  docTemplaterelation,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInforelation.InstanceName(), SwaggerInforelation)
}
