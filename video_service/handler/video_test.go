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

func TestCreateRecord(t *testing.T) {
	// 连接到数据库
	model.InitDb()

	// 创建一个 Video 实例，用于保存数据
	video := model.Videos{
		Common: model.Common{
			ID: 59549940993368064,
		},
		AuthID:        59469710022807552,
		PlayUrl:       "douyin.cfddfc.onlinedouyin/video/cfd--8a251c79-a0da-4166-aa1f-585e3616f931.mp4",
		CoverUrl:      "douyin.cfddfc.onlinedouyin/cover/cfd--8a251c79-a0da-4166-aa1f-585e3616f931.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "cfd",
	}

	// 使用 Create 方法插入数据
	result := model.DB.Create(&video)
	if result.Error != nil {
		// 处理错误
		fmt.Println(result.Error)
	} else {
		// 插入成功
		fmt.Println("Inserted video with ID:", video.ID)
	}
}
