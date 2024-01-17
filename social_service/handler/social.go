package handler

import (
	"golang.org/x/net/context"
	"social_service/model"
	"social_service/pkg/redis"
	social_pb "social_service/server"
	"utils/exception"
)

type SocialService struct {
	social_pb.UnimplementedSocialServiceServer
}

// FollowAction 点击关注
func (s *SocialService) FollowAction(ctx context.Context, req *social_pb.FollowRequest) (resp *social_pb.FollowResponse, err error) {
	resp = new(social_pb.FollowResponse)

	//初始化resp
	resp.StatusCode = exception.ERROR
	resp.StatusMsg = exception.GetMsg(exception.ERROR)

	//自己不能关注自己
	if req.UserId == req.ToUserId {
		resp.StatusCode = exception.FollowSelfErr
		resp.StatusMsg = exception.GetMsg(exception.FollowSelfErr)
		return resp, nil
	}

	err = redis.FollowAction(req.UserId, req.ToUserId, req.ActionType) //点赞操作，将信息存储在 redis 里面
	if err != nil {
		return resp, err
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}

// GetFollowList 关注列表
func (s *SocialService) GetFollowList(ctx context.Context, req *social_pb.FollowListRequest) (resp *social_pb.FollowListResponse, err error) {
	resp = new(social_pb.FollowListResponse)

	//初始化resp
	resp.StatusCode = exception.ERROR
	resp.StatusMsg = exception.GetMsg(exception.ERROR)

	err = redis.FollowList(req.UserId, &resp.UserId)
	if err != nil {
		return resp, err
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}

// GetFollowerList 粉丝列表（被关注列表）
func (s *SocialService) GetFollowerList(ctx context.Context, req *social_pb.FollowListRequest) (resp *social_pb.FollowListResponse, err error) {
	resp = new(social_pb.FollowListResponse)

	//初始化resp
	resp.StatusCode = exception.ERROR
	resp.StatusMsg = exception.GetMsg(exception.ERROR)

	err = redis.FollowerList(req.UserId, &resp.UserId)
	if err != nil {
		return resp, err
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}

// GetFriendList 好友列表
func (s *SocialService) GetFriendList(ctx context.Context, req *social_pb.FollowListRequest) (resp *social_pb.FollowListResponse, err error) {
	resp = new(social_pb.FollowListResponse)

	//初始化resp
	resp.StatusCode = exception.ERROR
	resp.StatusMsg = exception.GetMsg(exception.ERROR)

	err = redis.FriendList(req.UserId, &resp.UserId)
	if err != nil {
		return resp, err
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}

// PostMessage 发送信息
func (s *SocialService) PostMessage(ctx *context.Context, req *social_pb.PostMessageRequest) (resp *social_pb.PostMessageResponse, err error) {
	resp = new(social_pb.PostMessageResponse)
	message := model.Message{
		UserId:   req.UserId,
		ToUserId: req.ToUserId,
		Message:  req.Content,
	}
	err = model.GetMessageInstance().PostMessage(message)
	if err != nil {
		resp.StatusCode = exception.ERROR
		resp.StatusMsg = exception.GetMsg(exception.ERROR)
		return resp, nil
	}
	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}

// GetMessage 获取信息列表
func (s *SocialService) GetMessage(ctx context.Context, req *social_pb.GetMessageRequest) (resp *social_pb.GetMessageResponse, err error) {
	resp = new(social_pb.GetMessageResponse) // 结构体用new
	var messages []model.Message
	err = model.GetMessageInstance().GetMessage(req.UserId, req.ToUserId, req.PreMsgTime, &messages)
	if err != nil {
		resp.StatusCode = exception.ERROR
		resp.StatusMsg = exception.GetMsg(exception.ERROR)
		return resp, nil
	}

	for _, message := range messages {
		resp.Message = append(resp.Message, &social_pb.Message{
			Id:        message.Id,
			ToUserId:  message.ToUserId,
			UserId:    message.UserId,
			Content:   message.Message,
			CreatedAt: message.CreatedAt.Unix(),
		})
	}
	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}
