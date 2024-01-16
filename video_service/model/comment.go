package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"time"
	"utils/snowFlake"
)

type Comment struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    int64
	VideoId   int64
	Content   string `gorm:"default:(-)"` // 评论内容
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CommentModel struct {
}

// 懒汉单例模式
var commentModel *CommentModel

var commentOnce sync.Once

func GetCommentInstance() *CommentModel {
	commentOnce.Do(func() {
		commentModel = &CommentModel{}
	})
	return commentModel
}

// AddComment 新增评论
func (*CommentModel) AddComment(tx *gorm.DB, comment *Comment) (id int64, err error) {
	// 雪花算法分配唯一id
	flake, _ := snowFlake.NewSnowFlake(7, 2)
	comment.Id = flake.NextId()
	comment.CreatedAt = time.Now()

	res := tx.Create(&comment)
	if res.Error != nil {
		return -1, res.Error
	}
	return comment.Id, nil
}

// DeleteCommment 删除评论
func (*CommentModel) DeleteCommment(commentId int64) (err error) {
	var comment Comment
	res := DB.First(&comment, commentId)
	if res.Error != nil {
		return res.Error
	}
	if comment.DeletedAt.IsZero() {
		return errors.New("评论不存在")
	}

	comment.DeletedAt = time.Now()
	res = DB.Save(&comment)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetComment 根据评论id找到对应评论
func (*CommentModel) GetComment(tx *gorm.DB, commentId int64) (Comment, error) {
	comment := Comment{}
	res := tx.Model(&Comment{}).Where("id = ?", commentId).First(&comment)
	if res.Error != nil {
		return comment, res.Error
	}
	return comment, nil
}

func (*CommentModel) CommentList(videoId int64) (comments []Comment, err error) {
	comments = []Comment{}

	res := DB.Model(&Comment{}).Where("video_id = ?", videoId).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
