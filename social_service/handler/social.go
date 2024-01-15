package handler

import (
	"golang.org/x/net/context"
	"social_service/model"
	social_pb "social_service/server"
	utils "utils/status_code"
)

type SocialService struct {
	social_pb.UnimplementedSocialServiceServer
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
		resp.StatusCode = utils.ERROR
		resp.StatusMsg = utils.GetMsg(utils.ERROR)
		return resp, nil
	}
	resp.StatusCode = utils.SUCCESS
	resp.StatusMsg = utils.GetMsg(utils.SUCCESS)

	return resp, nil
}

// GetMessage 获取信息列表
func (s *SocialService) GetMessage(ctx context.Context, req *social_pb.GetMessageRequest) (resp *social_pb.GetMessageResponse, err error) {
	resp = new(social_pb.GetMessageResponse) // 结构体用new
	var messages []model.Message
	err = model.GetMessageInstance().GetMessage(req.UserId, req.ToUserId, req.PreMsgTime, &messages)
	if err != nil {
		resp.StatusCode = utils.ERROR
		resp.StatusMsg = utils.GetMsg(utils.ERROR)
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
	resp.StatusCode = utils.SUCCESS
	resp.StatusMsg = utils.GetMsg(utils.SUCCESS)
	return resp, nil
}
