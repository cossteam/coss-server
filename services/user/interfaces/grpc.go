package interfaces

import (
	"fmt"
	"im/services/user/api/v1"
	"im/services/user/domain/service"
	"im/services/user/utils"
)

type GrpcHandler struct {
	svc service.UserService
}

// 用户登录
func (g *GrpcHandler) UserLogin(request *api.UserLoginRequest) (*api.UserLoginResponse, error) {
	if err := g.validateLogin(request); err != nil {
		return nil, err
	}
	userInfo, err := g.svc.Login(request)
	if err != nil {
		return nil, err
	}
	//修改登录时间
	_, err = g.svc.UpdateUserInfo(userInfo)
	if err != nil {
		return nil, err
	}
	//参数校验
	return &api.UserLoginResponse{
		UserId:   userInfo.ID,
		NickName: userInfo.NickName,
		Email:    userInfo.Email,
		Tel:      userInfo.Tel,
		Avatar:   userInfo.Avatar,
	}, nil
}

// 用户注册
func (g *GrpcHandler) UserRegister(request *api.UserRegisterRequest) (*api.UserRegisterResponse, error) {
	//参数校验
	err := g.validateRegister(request)
	if err != nil {
		return nil, err
	}
	_, err = g.svc.GetUserInfoByEmail(request.Email)
	if err == nil {
		return nil, fmt.Errorf("邮箱已被注册")
	}
	//添加用户
	userInfo, err := g.svc.Register(request)
	return &api.UserRegisterResponse{
		UserId: userInfo.ID,
	}, nil

}

func (g *GrpcHandler) validateLogin(request *api.UserLoginRequest) error {
	if request.Email == "" || request.Password == "" {
		return fmt.Errorf("邮箱或密码不能为空")
	}

	//正则表达式判断邮箱是否正确
	if !utils.IsEmail(request.Email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	if !utils.ValidatePassword(request.Password) {
		return fmt.Errorf("密码格式不正确")
	}
	return nil
}
func (g *GrpcHandler) validateRegister(request *api.UserRegisterRequest) error {
	if request.Email == "" || request.Password == "" {
		return fmt.Errorf("邮箱或密码不能为空")
	}
	//正则表达式判断邮箱是否正确
	if !utils.IsEmail(request.Email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	if !utils.ValidatePassword(request.Password) || !utils.ValidatePassword(request.ConfirmPassword) {
		return fmt.Errorf("密码格式不正确")
	}
	if request.Password != request.ConfirmPassword {
		return fmt.Errorf("两次密码不一致")
	}
	return nil
}
