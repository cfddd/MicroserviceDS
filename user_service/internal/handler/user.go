package handler

import (
	"context"
	"user_service/internal/model"
	"user_service/server"
	exception "utils/status_code"
)

type UserService struct {
	server.UnimplementedUserServiceServer // 版本兼容问题
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserRefister(ctx context.Context, req *server.UserRequest) (resp *server.UserResponse, err error) {
	resp = new(server.UserResponse)
	var user model.User

	//检查用户是否存在
	exist := model.GetInstance().CheckUserExist(req.Username)
	if exist == false {
		resp.StatusCode = exception.UserExist
		resp.StatusMsg = exception.GetMsg(exception.UserExist)
		return resp, err
	}

	user.Name = req.Username
	user.Password = req.Password

	//创建用户
	err = model.GetInstance().Create(&user)
	if err != nil {
		resp.StatusCode = exception.DataErr
		resp.StatusMsg = exception.GetMsg(exception.DataErr)
		return resp, nil
	}

	//查询用户ID
	userName, err := model.GetInstance().FindUserByName(user.Name)
	if err != nil {
		resp.StatusCode = exception.UserUnExist
		resp.StatusMsg = exception.GetMsg(exception.UserUnExist)
		return resp, err
	} else {
		resp.StatusCode = exception.SUCCESS
		resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
		resp.UserId = userName.ID
		return resp, nil
	}
}
