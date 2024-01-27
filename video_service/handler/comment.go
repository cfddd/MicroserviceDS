package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
	"video_service/model"
	"video_service/pkg/cache"
	video_pb "video_service/server"
	utils "video_service/utils/exception"
)

func (v *VideoService) CommentAction(ctx context.Context, req *video_pb.CommentActionRequest) (resp *video_pb.CommentActionResponse, err error) {
	resp = new(video_pb.CommentActionResponse)
	key := fmt.Sprintf("video:comment_list:%s", strconv.FormatInt(req.VideoId, 10))
	comment := model.Comments{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Content: req.CommentText,
	}

	time := time.Now()
	if req.ActionType == 1 { // 发布评论
		comment.CreatedAt = time

		// 在事务中执行创建操作
		tx := model.DB.Begin()
		// 因为后续评论要存入缓存中
		id, err := model.GetCommentInstance().AddComment(tx, &comment)
		if err != nil {
			resp.StatusCode = utils.CommentErr
			resp.StatusMsg = utils.GetMsg(utils.CommentErr)
			resp.Comment = nil
			return resp, err
		}

		comment.Id = id

		commentJson, _ := json.Marshal(comment)
		// 存入缓存
		member := redis.Z{
			Score:  float64(time.Unix()),
			Member: commentJson,
		}

		err = cache.Redis.ZAdd(cache.Ctx, key, &member).Err()
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
		tx.Commit()

		commentResp := &video_pb.Comment{
			Id:        comment.Id,
			Content:   comment.Content,
			CreatedAt: time.Format("01-02"),
		}

		resp.StatusCode = utils.SUCCESS
		resp.StatusMsg = utils.GetMsg(utils.SUCCESS)
		resp.Comment = commentResp

		return resp, nil
	} else { // 删除评论
		// 事务
		tx := model.DB.Begin()
		commentInstance, err := model.GetCommentInstance().GetComment(tx, req.CommentId)

		commentJson, _ := json.Marshal(commentInstance)

		err = model.GetCommentInstance().DeleteCommment(req.CommentId)
		if err != nil {
			resp.StatusCode = utils.CommentDeleteErr
			resp.StatusMsg = utils.GetMsg(utils.CommentDeleteErr)
			resp.Comment = nil

			return resp, err
		}

		log.Print(commentJson)

		// 删除缓存
		count, err := cache.Redis.ZRem(cache.Ctx, key, string(commentJson)).Result()
		log.Printf("删除了: %v 条评论", count)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("缓存错误,%v", err)
		}

		tx.Commit()

		resp.StatusCode = utils.SUCCESS
		resp.StatusMsg = utils.GetMsg(utils.SUCCESS)
		resp.Comment = nil
		return resp, nil
	}
}

// CommentList 获取评论列表
func (*VideoService) CommentList(ctx context.Context, req *video_pb.CommentListRequest) (resp *video_pb.CommentListResponse, err error) {
	resp = new(video_pb.CommentListResponse)
	var comments []model.Comments
	key := fmt.Sprintf("video:comment_list:%s", strconv.FormatInt(req.VideoId, 10))

	count, err := cache.Redis.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Print(err)
		return resp, fmt.Errorf("读取缓存错误,%v", err)
	}

	if count == 0 {
		err := buildCommentCache(req.VideoId)
		if err != nil {
			log.Printf("构建缓存错误,%v", err)
			return resp, fmt.Errorf("构建缓存错误,%v", err)
		}
	}

	// 查询缓存
	commentsString, err := cache.Redis.ZRevRange(cache.Ctx, key, 0, -1).Result()
	if err != nil {
		log.Printf("读取缓存错误,%v", err)
		return resp, fmt.Errorf("读取缓存错误,%v", err)
	}

	for _, commentString := range commentsString {
		var comment model.Comments
		err := json.Unmarshal([]byte(commentString), &comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	resp.StatusCode = utils.SUCCESS
	resp.StatusMsg = utils.GetMsg(utils.SUCCESS)
	resp.CommentList = BuildComments(comments)

	return resp, nil
}

// BuildComments 转格式
func BuildComments(comments []model.Comments) []*video_pb.Comment {
	var commentResp []*video_pb.Comment
	for _, comment := range comments {
		commentResp = append(commentResp, &video_pb.Comment{
			Id:        comment.Id,
			UserId:    comment.UserId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("01-02"),
		})
	}
	return commentResp
}

// buildCommentCache 构建评论列表缓存
func buildCommentCache(videoId int64) error {
	key := fmt.Sprintf("video:comment_list:%s", strconv.FormatInt(videoId, 10))

	comments, err := model.GetCommentInstance().CommentList(videoId)
	if err != nil {
		return err
	}

	var members []*redis.Z
	for _, comment := range comments {
		commentJson, err := json.Marshal(comment)
		if err != nil {
			log.Printf("构建缓存错误,%v", err)
			continue
		}
		members = append(members, &redis.Z{
			Score:  float64(comment.CreatedAt.Unix()), //根据评论时间的先后对缓存中的评论进行排序
			Member: commentJson,
		})
	}

	err = cache.Redis.ZAdd(cache.Ctx, key, members...).Err()
	return err
}

// GetCommentCount 通过缓存查看视频得评论数量
func GetCommentCount(videoId int64) int64 {
	key := fmt.Sprintf("video:comment_list:%s", strconv.FormatInt(videoId, 10))

	exists, err := cache.Redis.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Print(err)
	}

	if exists == 0 { //如果缓存不存在就构建缓存
		err := buildCommentCache(videoId)
		if err != nil {
			log.Print(err)
		}
	}

	count, err := cache.Redis.ZCard(cache.Ctx, key).Result()
	if err != nil {
		log.Print(err)
	}

	return count
}
