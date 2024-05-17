// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package v1

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for UserInfoStatus.
const (
	N0 UserInfoStatus = 0
	N1 UserInfoStatus = 1
	N2 UserInfoStatus = 2
	N3 UserInfoStatus = 3
	N4 UserInfoStatus = 4
)

// Preferences defines model for Preferences.
type Preferences struct {
	// OpenBurnAfterReading 是否开启阅后即焚
	OpenBurnAfterReading        bool   `json:"open_burn_after_reading"`
	OpenBurnAfterReadingTimeOut int    `json:"open_burn_after_reading_time_out"`
	Remark                      string `json:"remark"`

	// SilentNotification 是否开启静默通知
	SilentNotification bool `json:"silent_notification"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	Avatar         string       `json:"avatar"`
	CossId         string       `json:"coss_id"`
	Email          string       `json:"email"`
	LastLoginTime  int64        `json:"last_login_time,omitempty"`
	LoginNumber    int64        `json:"login_number,omitempty"`
	NewDeviceLogin bool         `json:"new_device_login,omitempty"`
	Nickname       string       `json:"nickname"`
	Preferences    *Preferences `json:"preferences,omitempty"`
	RelationStatus int          `json:"relation_status"`
	Signature      string       `json:"signature"`

	// Status 0: 未知状态 1: 正常状态 2: 被禁用 3: 已删除 4: 锁定状态
	Status UserInfoStatus `json:"status"`
	Tel    string         `json:"tel"`
	UserId string         `json:"user_id"`
}

// UserInfoStatus 0: 未知状态 1: 正常状态 2: 被禁用 3: 已删除 4: 锁定状态
type UserInfoStatus int

// UserLoginClient defines model for UserLoginClient.
type UserLoginClient struct {
	// ClientIp 客户端ip
	ClientIp string `json:"client_ip"`

	// DriverId 设备id
	DriverId string `json:"driver_id"`

	// DriverType 设备类型
	DriverType string `json:"driver_type"`

	// LoginAt 登录时间
	LoginAt int64 `json:"login_at"`
}

// UserSecretBundle defines model for UserSecretBundle.
type UserSecretBundle struct {
	// SecretBundle 用户密钥包
	SecretBundle string `json:"secret_bundle"`

	// UserId 用户id
	UserId string `json:"user_id"`
}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody struct {
	// Avatar 用户头像
	Avatar string `json:"avatar"`
	CossId string `json:"coss_id"`

	// Nickname 用户昵称
	Nickname string `json:"nickname"`

	// Signature 个性签名
	Signature string `json:"signature"`

	// Tel 手机号
	Tel string `json:"tel"`
}

// UserActivateParams defines parameters for UserActivate.
type UserActivateParams struct {
	// UserId 用户ID
	UserId string `form:"user_id" json:"user_id"`

	// Key 激活密钥
	Key string `form:"key" json:"key"`
}

// UpdateUserAvatarMultipartBody defines parameters for UpdateUserAvatar.
type UpdateUserAvatarMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

// UpdateUserBundleJSONBody defines parameters for UpdateUserBundle.
type UpdateUserBundleJSONBody struct {
	// SecretBundle 用户密钥包
	SecretBundle string `json:"secret_bundle"`
}

// UserEmailVerificationJSONBody defines parameters for UserEmailVerification.
type UserEmailVerificationJSONBody struct {
	Email openapi_types.Email `json:"email"`
}

// UserLoginJSONBody defines parameters for UserLogin.
type UserLoginJSONBody struct {
	// DriverId 当前登录设备的唯一标识符
	DriverId string `json:"driver_id"`

	// DriverToken 当前设备的设备token 用于推送手机端的系统通知
	DriverToken string `json:"driver_token"`

	// Email 用户的电子邮件地址
	Email string `json:"email"`

	// Password 用户密码
	Password string `json:"password"`

	// Platform 用户登录的平台(ios、android、web、huawei...)
	Platform string `json:"platform"`
}

// UserLogoutJSONBody defines parameters for UserLogout.
type UserLogoutJSONBody struct {
	DriverId string `json:"driver_id"`
}

// UpdateUserPasswordJSONBody defines parameters for UpdateUserPassword.
type UpdateUserPasswordJSONBody struct {
	// ConfirmPassword 确认密码
	ConfirmPassword string `json:"confirm_password"`

	// OldPassword 旧密码
	OldPassword string `json:"old_password"`

	// Password 新密码
	Password string `json:"password"`
}

// SetUserPublicKeyJSONBody defines parameters for SetUserPublicKey.
type SetUserPublicKeyJSONBody struct {
	PublicKey string `json:"public_key"`
}

// ResetUserPublicKeyJSONBody defines parameters for ResetUserPublicKey.
type ResetUserPublicKeyJSONBody struct {
	Code      string               `json:"code"`
	Email     *openapi_types.Email `json:"email,omitempty"`
	PublicKey string               `json:"public_key"`
}

// UserRegisterJSONBody defines parameters for UserRegister.
type UserRegisterJSONBody struct {
	// ConfirmPassword 确认密码
	ConfirmPassword string `json:"confirm_password"`

	// Email 用户邮箱
	Email openapi_types.Email `json:"email"`

	// Nickname 用户昵称
	Nickname string `json:"nickname"`

	// Password 用户密码
	Password string `json:"password"`

	// PublicKey 用户公钥
	PublicKey string `json:"public_key"`
}

// SearchUserParams defines parameters for SearchUser.
type SearchUserParams struct {
	// Email 用户邮箱
	Email string `form:"email" json:"email"`
}

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// UpdateUserAvatarMultipartRequestBody defines body for UpdateUserAvatar for multipart/form-data ContentType.
type UpdateUserAvatarMultipartRequestBody UpdateUserAvatarMultipartBody

// UpdateUserBundleJSONRequestBody defines body for UpdateUserBundle for application/json ContentType.
type UpdateUserBundleJSONRequestBody UpdateUserBundleJSONBody

// UserEmailVerificationJSONRequestBody defines body for UserEmailVerification for application/json ContentType.
type UserEmailVerificationJSONRequestBody UserEmailVerificationJSONBody

// UserLoginJSONRequestBody defines body for UserLogin for application/json ContentType.
type UserLoginJSONRequestBody UserLoginJSONBody

// UserLogoutJSONRequestBody defines body for UserLogout for application/json ContentType.
type UserLogoutJSONRequestBody UserLogoutJSONBody

// UpdateUserPasswordJSONRequestBody defines body for UpdateUserPassword for application/json ContentType.
type UpdateUserPasswordJSONRequestBody UpdateUserPasswordJSONBody

// SetUserPublicKeyJSONRequestBody defines body for SetUserPublicKey for application/json ContentType.
type SetUserPublicKeyJSONRequestBody SetUserPublicKeyJSONBody

// ResetUserPublicKeyJSONRequestBody defines body for ResetUserPublicKey for application/json ContentType.
type ResetUserPublicKeyJSONRequestBody ResetUserPublicKeyJSONBody

// UserRegisterJSONRequestBody defines body for UserRegister for application/json ContentType.
type UserRegisterJSONRequestBody UserRegisterJSONBody