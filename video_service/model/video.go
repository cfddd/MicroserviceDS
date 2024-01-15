package model

import (
	"gorm.io/gorm"
	"time"
)

type common struct {
	ID        uint64         `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Video struct {
	common
	VideoCreator  int64  `gorm:"column:video_creator;type:bigint(20)" json:"video_creator"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(256)" json:"play_url"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(256)" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:bigint(20)" json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count;type:bigint(20)" json:"comment_count"`
	Title         string `gorm:"column:title;type:varchar(256)" json:"title"`
}
