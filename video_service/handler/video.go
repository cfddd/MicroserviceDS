package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"sync"
	"time"
	"utils/exception"
	"video_service/logger"
	"video_service/model"
	"video_service/pkg/cache"
	"video_service/pkg/cut"
	"video_service/pkg/db"
	"video_service/pkg/oss7"
	video_server "video_service/server"
)

type VideoService struct {
	video_server.UnimplementedVideoServiceServer // 版本兼容问题
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (v *VideoService) Feed(ctx context.Context, req *video_server.FeedRequest) (resp *video_server.FeedResponse, err error) {
	resp = new(video_server.FeedResponse)

	//初始化
	resp.StatusCode = exception.VideoUnExist
	resp.StatusMsg = exception.GetMsg(exception.VideoUnExist)

	//拿到时间
	var posTime time.Time
	if req.LatestTime == -1 {
		posTime = time.Now()
	} else {
		posTime = time.Unix(req.LatestTime/1000, 0)
	}

	videoList, err := model.GetVideoInstance().GetVideoByTime(posTime)
	if err != nil {
		return resp, err
	}

	resp.VideoList = model.BuildVideo(videoList, req.UserId)

	// 获取列表中最早发布视频的时间作为下一次请求的时间
	LastIndex := len(videoList) - 1
	resp.NextTime = videoList[LastIndex].CreatedAt.Unix()

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}

// PublishAction 发布视频
func (*VideoService) PublishAction(ctx context.Context, req *video_server.PublishActionRequest) (resp *video_server.PublishActionResponse, err error) {
	resp = new(video_server.PublishActionResponse)
	reqString, err := json.Marshal(&req)

	// 放入消息队列
	conn := db.InitMQ()
	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Print(err)
	}
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"video_publish",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}
	// 5s 后，如果没有消费则自动删除
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{ // 发送的消息，固定有消息体和一些额外的消息头，包中提供了封装对象
			ContentType: "application/octet-stream",
			Body:        reqString, // 请求信息重新封装为json，加入消息队列
		})
	if err != nil {
		resp.StatusCode = exception.VideoUploadErr
		resp.StatusMsg = exception.GetMsg(exception.VideoUploadErr)

		return resp, nil
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}

// PublishAction1 发布视频
func (*VideoService) PublishAction1(ctx context.Context, req *video_server.PublishActionRequest) (resp *video_server.PublishActionResponse, err error) {
	var updataErr, creatErr error
	resp = new(video_server.PublishActionResponse)
	key := fmt.Sprintf("%s:%s", "user", "work_count")

	// 获取参数,生成地址
	title := req.Title
	UUID := uuid.New()
	videoDir := "douyin/video/" + title + "--" + UUID.String() + ".mp4"
	pictureDir := "douyin/cover/" + title + "--" + UUID.String() + ".jpg"

	videoUrl := viper.GetString("oss.link") + videoDir
	pictureUrl := viper.GetString("oss.link") + pictureDir

	// 等待上传和创建数组库完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 上传视频，切取封面，上传图片
	go func() {
		defer wg.Done()
		// 上传视频
		updataErr = oss7.UploadFileWithByte(videoDir, req.Data)
		// 获取封面,获取第几秒的封面
		coverByte, _ := cut.Cover(videoUrl, "1.0")
		// 上传封面
		updataErr = oss7.UploadFileWithByte(pictureDir, coverByte)
		logger.Log.Info("上传成功")
	}()

	// 创建数据
	go func() {
		defer wg.Done()
		// 创建video
		// CreatedAt ? 不让写，会报错
		video := model.Video{
			VideoCreator: req.UserId,
			Title:        title,
			CoverUrl:     pictureUrl,
			PlayUrl:      videoUrl,
		}
		creatErr = model.GetVideoInstance().Create(&video)
	}()

	wg.Wait()

	// 异步回滚
	if updataErr != nil || creatErr != nil {
		go func() {
			// 存入数据库失败，删除上传
			if creatErr != nil {
				_ = oss7.DeleteFile(videoDir)
				_ = oss7.DeleteFile(pictureDir)
			}
			// 上传失败，删除数据库
			if updataErr != nil {
				// TODO 根据url查找，效率比较低
				_ = model.GetVideoInstance().DeleteVideoByUrl(videoUrl)
			}
		}()
	}
	if updataErr != nil || creatErr != nil {
		resp.StatusCode = exception.VideoUploadErr
		resp.StatusMsg = exception.GetMsg(exception.VideoUploadErr)
		return resp, updataErr
	}

	// 发布成功，缓存中作品总数 + 1，如果不存在缓存则不做操作
	exist, err := cache.Redis.HExists(cache.Ctx, key, strconv.FormatInt(req.UserId, 10)).Result()
	if err != nil {
		return nil, fmt.Errorf("缓存错误：%v", err)
	}

	if exist {
		// 字段存在，该记录数量 + 1
		_, err = cache.Redis.HIncrBy(cache.Ctx, key, strconv.FormatInt(req.UserId, 10), 1).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
	}

	// 发布成功延时双删发布列表
	workKey := fmt.Sprintf("%s:%s:%s", "user", "work_list", strconv.FormatInt(req.UserId, 10))
	err = cache.Redis.Del(cache.Ctx, workKey).Err()
	if err != nil {
		return nil, fmt.Errorf("缓存错误：%v", err)
	}
	defer func() {
		go func() {
			//延时3秒执行
			time.Sleep(time.Second * 3)
			//再次删除缓存
			cache.Redis.Del(cache.Ctx, workKey)
		}()
	}()

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}
