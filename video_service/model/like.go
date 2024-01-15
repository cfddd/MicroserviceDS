package model

type Like struct {
	common
	UserId  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`
	VideoId uint64 `gorm:"column:video_id;type:bigint(20) unsigned;NOT NULL" json:"video_id"`
}
