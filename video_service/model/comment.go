package model

type Comment struct {
	common
	VideoId uint64 `gorm:"column:video_id;type:bigint(20) unsigned" json:"video_id"`
	UserId  int64  `gorm:"column:user_id;type:bigint(20)" json:"user_id"`
	Content string `gorm:"column:content;type:varchar(2048)" json:"content"`
}
