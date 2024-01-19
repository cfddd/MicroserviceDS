package model

import (
	"sync"
	"time"
	"utils/snowFlake"
)

type Messages struct {
	Id         int64     `json:"id" gorm:"primary_key"`
	UserId     int64     `json:"user_id" gorm:"foreignKey:UserId;column:user_id"` // 指定外键关系
	FollowToID int64     `json:"follow_to_id" gorm:"column:follow_to_id"`
	Content    string    `json:"content" gorm:"column:content"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}

type MessageModel struct {
}

var messageModel *MessageModel
var messageOnce sync.Once

// GetMessageInstance 获取单例实例
func GetMessageInstance() *MessageModel {
	messageOnce.Do(func() {
		messageModel = &MessageModel{}
	})
	return messageModel
}

func (m *MessageModel) PostMessage(message Messages) error {
	message.CreatedAt = time.Now()
	// 雪花算法生成唯一id
	flake, _ := snowFlake.NewSnowFlake(7, 3)
	message.Id = flake.NextId()
	err := DB.Create(&message).Error
	return err
}

func (m *MessageModel) GetMessage(uId, tId, preMsgTime int64, messages *[]Messages) error {
	time := time.Unix(preMsgTime, 0)
	err := DB.Model(&Messages{}).Where("user_id + follow_to_id = ? and created_at > ? ", uId+tId, time).Find(messages).Error
	return err
}
