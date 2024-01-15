package handler

import (
	"github.com/gin-gonic/gin"
	video_pb "video_service/server"
)

func (v *VideoService) CommentAction(ctx *gin.Context, req *video_pb.CommentActionRequest) (resp *video_pb.CommentActionResponse, err error) {
	resp = new(video_pb.CommentActionResponse)
	// 判断是删除评论还是增加评论

	comment :=
	return resp, nil
}
