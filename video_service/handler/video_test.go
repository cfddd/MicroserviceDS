package handler

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
	"time"
	"video_service/config"
	"video_service/model"
)

func TestCommon(t *testing.T) {
	cfd := model.Videos{
		Common: model.Common{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Title: "cfd",

		AuthID:        123456,
		VideoCreator:  789012,
		PlayUrl:       "http://example.com/play",
		CoverUrl:      "http://example.com/cover",
		FavoriteCount: 100,
		CommentCount:  50,
	}
	fmt.Println(cfd.ID)
}

func TestRabbitMQ(t *testing.T) {
	config.InitConfig()
	// 连接到RabbitMQ服务器
	url := config.InitRabbitMQUrl()
	t.Log(url)
	_, err := amqp.Dial(url)
	if err != nil {
		t.Errorf("RabbitMQ 连接失败！%v", err)
		return
	}
	t.Log("RabbitMQ 连接成功!")
}
