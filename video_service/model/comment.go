package model

import "time"

type Comment struct {
	Id            int64 `gorm:"primaryKey"`
	UserId        int64
	VideoId       int64
	Content       string `gorm:"default:(-)"` // 评论内容
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	CommentStatus bool `gorm:"default:(-)"`
}
