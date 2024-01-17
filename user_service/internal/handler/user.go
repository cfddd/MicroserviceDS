package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"user_service/internal/model"
	"user_service/pkg/redis"
	"user_service/server"
	"utils/exception"
)

type UserService struct {
	server.UnimplementedUserServiceServer // 版本兼容问题
}

func NewUserService() *UserService {
	return &UserService{}
}

// 用户注册
func (*UserService) UserRegister(ctx context.Context, req *server.UserRequest) (resp *server.UserResponse, err error) {
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

	//根据用户名查询用户ID
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

// 用户登录
func (*UserService) UserLogin(ctx context.Context, req *server.UserRequest) (resp *server.UserResponse, err error) {
	resp = new(server.UserResponse)

	//检查用户是否存在
	exist := model.GetInstance().CheckUserExist(req.Username)
	if exist {
		resp.StatusCode = exception.UserUnExist
		resp.StatusMsg = exception.GetMsg(exception.UserUnExist)
		return resp, err
	}

	//检查用户密码是否正确
	user, err := model.GetInstance().FindUserByName(req.Username)
	ok := model.GetInstance().CheckPassWord(req.Password, user.Password)
	if !ok {
		resp.StatusCode = exception.PasswordError
		resp.StatusMsg = exception.GetMsg(exception.PasswordError)
		return resp, err
	} else {
		resp.StatusCode = exception.SUCCESS
		resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
		return resp, nil
	}
}

func RespUser(u *model.User) *server.User {
	user := server.User{
		Id:   u.ID,
		Name: u.Name,
	}
	return &user
}

// 查询用户信息
func (*UserService) UserInfo(ctx context.Context, req *server.UserInfoRequest) (resp *server.UserInfoResponse, err error) {
	resp = new(server.UserInfoResponse)

	//根据userId切片查询用户信息
	userIds := req.UserIds

	for _, userId := range userIds {
		//查看缓存是否存在，若存在获取需要的信息
		var user *model.User
		key := fmt.Sprintf("%s:%s:%s", "user", "info", strconv.FormatInt(userId, 10))

		exists, err := redis.Redis.Exists(redis.Ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误: %v", err)
		}

		//exists>0表示有缓存
		if exists > 0 {
			//查询缓存
			userString, err := redis.Redis.Get(redis.Ctx, key).Result()
			if err != nil {
				return nil, fmt.Errorf("缓存错误: %v", err)
			}
			err = json.Unmarshal([]byte(userString), &user)
			if err != nil {
				return nil, err
			}
		} else {
			//查询数据库
			user, err = model.GetInstance().FindUserById(userId)
			if err != nil {
				resp.StatusCode = exception.UserUnExist
				resp.StatusMsg = exception.GetMsg(exception.UserUnExist)
				return resp, err
			}

			//将查询结果放入缓存中
			userJson, _ := json.Marshal(&user)
			err = redis.Redis.Set(redis.Ctx, key, userJson, 24*time.Hour).Err()
			if err != nil {
				return nil, fmt.Errorf("缓存错误: %v", err)
			}
		}
		resp.Users = append(resp.Users, RespUser(user))
	}
	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}
