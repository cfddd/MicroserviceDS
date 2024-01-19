package handler

import (
	"golang.org/x/net/context"
	"log"
	"social_service/model"
	"social_service/pkg/redis"
	social_pb "social_service/server"
	"utils/exception"
)

type SocialService struct {
	social_pb.UnimplementedSocialServiceServer
}

func NewSocialService() *SocialService {
	return &SocialService{}
}

// FollowAction 点击关注
func (s *SocialService) FollowAction(ctx context.Context, req *social_pb.FollowRequest) (resp *social_pb.FollowResponse, err error) {
	resp = new(social_pb.FollowResponse)

	//初始化resp
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
func (s *SocialService) PostMessage(ctx context.Context, req *social_pb.PostMessageRequest) (resp *social_pb.PostMessageResponse, err error) {
	resp = new(social_pb.PostMessageResponse)
	message := model.Messages{
		UserId:     req.UserId,
		FollowToID: req.ToUserId,
		Content:    req.Content,
	}
	err = model.GetMessageInstance().PostMessage(message)
	if err != nil {
		resp.StatusCode = exception.ERROR
		resp.StatusMsg = exception.GetMsg(exception.ERROR)
		log.Printf("PostMessage: %v", err)
		return resp, nil
	}
	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}

// GetMessage 获取信息列表
func (s *SocialService) GetMessage(ctx context.Context, req *social_pb.GetMessageRequest) (resp *social_pb.GetMessageResponse, err error) {
	resp = new(social_pb.GetMessageResponse) // 结构体用new
	var messages []model.Messages
	err = model.GetMessageInstance().GetMessage(req.UserId, req.ToUserId, req.PreMsgTime, &messages)
	if err != nil {
		resp.StatusCode = exception.ERROR
		resp.StatusMsg = exception.GetMsg(exception.ERROR)
		return resp, nil
	}

	for _, message := range messages {
		resp.Message = append(resp.Message, &social_pb.Message{
			Id:         message.Id,
			FollowToId: message.FollowToID,
			UserId:     message.UserId,
			Content:    message.Content,
			CreatedAt:  message.CreatedAt.Unix(),
		})
	}
	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	return resp, nil
}
func (*SocialService) GetFollowInfo(ctx context.Context, req *social_pb.FollowInfoRequest) (resp *social_pb.FollowInfoResponse, err error) {
	resp = new(social_pb.FollowInfoResponse)
	for _, toUserId := range req.ToUserId {
		/* mysql
		res1, err1 := model.GetFollowInstance().IsFollow(req.UserId, toUserId)
		cnt2, err2 := model.GetFollowInstance().GetFollowCount(toUserId)
		cnt3, err3 := model.GetFollowInstance().GetFollowerCount(toUserId)
		*/
		res1, err1 := redis.IsFollow(req.UserId, toUserId)
		cnt2, err2 := redis.FollowCount(toUserId)
		cnt3, err3 := redis.FollowerCount(toUserId)
		if err1 != nil || err2 != nil || err3 != nil {
			resp.StatusCode = exception.ERROR
			resp.StatusMsg = exception.GetMsg(exception.ERROR)
			return resp, nil
		}
		resp.FollowInfo = append(resp.FollowInfo, &social_pb.FollowInfo{
			IsFollow:      res1,
			FollowCount:   cnt2,
			FollowerCount: cnt3,
			ToUserId:      toUserId,
		})
	}
	return resp, nil
}
