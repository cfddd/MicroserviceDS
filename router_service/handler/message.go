package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	res "router_service/response"
	service "router_service/server"
	"strconv"
	"utils/exception"
)

func PostMessage(ctx *gin.Context) {
	var postMessage service.PostMessageRequest
	userId, _ := ctx.Get("user_id")
	postMessage.UserId, _ = userId.(int64)
	toUserId := ctx.Query("to_user_id")
	postMessage.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	actionType := ctx.Query("action_type")
	actionTypeInt64, _ := strconv.ParseInt(actionType, 10, 32)

	if actionTypeInt64 != 1 {
		r := res.FavoriteActionResponse{
			StatusCode: exception.ErrOperate,
			StatusMsg:  exception.GetMsg(exception.ErrOperate),
		}

		ctx.JSON(http.StatusOK, r)
		return
	}

	postMessage.ActionType = int32(actionTypeInt64)
	content := ctx.Query("content")
	postMessage.Content = content

	socialServiceClient := ctx.Keys["social_service"].(service.SocialServiceClient)
	socialResp, err := socialServiceClient.PostMessage(context.Background(), &postMessage)
	if err != nil {
		PanicIfMessageError(err)
	}

	r := res.PostMessageResponse{
		StatusCode: socialResp.StatusCode,
		StatusMsg:  socialResp.StatusMsg,
	}
	ctx.JSON(http.StatusOK, r)
}

func GetMessage(ctx *gin.Context) {
	var getMessage service.GetMessageRequest
	userId, _ := ctx.Get("user_id")
	getMessage.UserId, _ = userId.(int64)
	toUserId := ctx.Query("to_user_id")
	getMessage.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	PreMsgTime := ctx.Query("pre_msg_time")
	getMessage.PreMsgTime, _ = strconv.ParseInt(PreMsgTime, 10, 64)

	socialServiceClient := ctx.Keys["social_service"].(service.SocialServiceClient)
	socialResp, err := socialServiceClient.GetMessage(context.Background(), &getMessage)

	if err != nil {
		PanicIfMessageError(err)
	}

	r := new(res.GetMessageResponse)
	r.StatusCode = socialResp.StatusCode
	r.StatusMsg = socialResp.StatusMsg
	for _, message := range socialResp.Message {
		messageResp := res.Message{
			Id:         message.Id,
			ToUserId:   message.ToUserId,
			FromUserID: message.UserId,
			Content:    message.Content,
			CreateTime: message.CreatedAt,
		}
		r.MessageList = append(r.MessageList, messageResp)
	}

	ctx.JSON(http.StatusOK, r)
}
