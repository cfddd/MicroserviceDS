package handler

import (
	"github.com/pkg/errors"
	"router_service/logger"
)

// PanicIfVideoError 视频错误处理
func PanicIfVideoError(err error) {
	if err != nil {
		err = errors.New("videoService--error--" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}

// PanicIfUserError 用户错误处理
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("UserService--error" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}

// PanicIfMessageError 消息错误处理
func PanicIfMessageError(err error) {
	if err != nil {
		err = errors.New("MessageService--error" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}

// PanicIfFollowError 关注错误处理
func PanicIfFollowError(err error) {
	if err != nil {
		err = errors.New("FollowService--error" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}

// PanicIfFavoriteError 喜欢错误处理
func PanicIfFavoriteError(err error) {
	if err != nil {
		err = errors.New("FavoriteService--error" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}

// PanicIfCommentError 评论错误处理
func PanicIfCommentError(err error) {
	if err != nil {
		err = errors.New("CommentService--error" + err.Error())
		logger.Log.Info(err)
		panic(err)
	}
}
