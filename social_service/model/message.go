package model

import (
	"sync"
	"time"
	"utils/snowFlake"
)

type Messages struct {
	Id        int64 `gorm:"primary_key"`
	UserId    int64
	ToUserId  int64
	Message   string
	CreatedAt time.Time
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
	err := DB.Model(&Messages{}).Where("UserId + ToUserId = ? and CreatedAt > ? ", uId+tId, preMsgTime).Find(messages).Error
	return err
}
