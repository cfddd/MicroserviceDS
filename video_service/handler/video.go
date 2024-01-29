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
	"video_service/logger"
	"video_service/model"
	"video_service/pkg/cache"
	"video_service/pkg/cut"
	"video_service/pkg/oss7"
	"video_service/pkg/rabbitMq"
	video_server "video_service/server"
	"video_service/utils/exception"
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

	resp.VideoList = BuildVideo(videoList, req.UserId)

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
	conn := rabbitMq.InitMQ()
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

func PublishVideo() {
	// 放入消息队列
	conn := rabbitMq.InitMQ()
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

	// 消费者
	msgs, err := ch.Consume(
		q.Name,
		"video_service",
		false, //手动确认
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}
	var forever chan struct{}
	go func() {
		logger.Log.Info("开始")
		for d := range msgs {
			logger.Log.Info("开始消费消息")
			var req video_server.PublishActionRequest
			err := json.Unmarshal(d.Body, &req)
			if err != nil {
				logger.Log.Error(err)
			}

			var updataErr, creatErr error
			key := fmt.Sprintf("%s:%s", "user", "work_count")

			// 获取参数,生成地址
			title := req.Title
			UUID := uuid.New()
			videoDir := "douyin/video/" + title + "--" + UUID.String() + ".mp4"
			pictureDir := "douyin/cover/" + title + "--" + UUID.String() + ".jpg"

			videoUrl := "http://" + viper.GetString("oss.link") + "/" + videoDir
			pictureUrl := "http://" + viper.GetString("oss.link") + "/" + pictureDir

			// 等待上传和创建数组库完成
			var wg sync.WaitGroup
			wg.Add(2)

			// 上传视频，切取封面，上传图片
			go func() {
				defer wg.Done()
				// 上传视频
				updataErr = oss7.UploadFileWithByte(videoDir, req.Data)
				// 获取封面,获取第几秒的封面
				coverByte, err := cut.Cover(videoUrl, "00:00:01")
				if err != nil {
					logger.Log.Error(err)
				}
				// 上传封面
				updataErr = oss7.UploadFileWithByte(pictureDir, coverByte)
				log.Print("上传成功")
			}()

			var videoID int64
			// 创建数据
			go func() {
				defer wg.Done()
				// 创建video
				// CreatedAt ? 不让写，会报错
				video := model.Videos{
					AuthID:   req.UserId,
					Title:    title,
					CoverUrl: pictureUrl,
					PlayUrl:  videoUrl,
				}
				videoID, creatErr = model.GetVideoInstance().Create(video)
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
						// 根据url查找，效率比较低
						// 使用id查找
						_ = model.GetVideoInstance().DeleteVideoByID(videoID)
					}
				}()
			}

			d.Ack(false) // 手动确认消息

			// 发布成功，缓存中作品总数 + 1，如果不存在缓存则不做操作
			exist, err := cache.Redis.HExists(cache.Ctx, key, strconv.FormatInt(req.UserId, 10)).Result()
			if err != nil {
				log.Print(err)
			}

			if exist {
				// 字段存在，该记录数量 + 1
				_, err = cache.Redis.HIncrBy(cache.Ctx, key, strconv.FormatInt(req.UserId, 10), 1).Result()
				if err != nil {
					log.Print(err)
				}
			}

			// 发布成功延时双删发布列表
			workKey := fmt.Sprintf("%s:%s:%s", "user", "work_list", strconv.FormatInt(req.UserId, 10))
			err = cache.Redis.Del(cache.Ctx, workKey).Err()
			if err != nil {
				log.Print(err)
			}

			go func() {
				//延时3秒执行
				time.Sleep(time.Second * 3)
				//再次删除缓存
				cache.Redis.Del(cache.Ctx, workKey)
			}()
		}
	}()
	<-forever
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

	videoUrl := "http://" + viper.GetString("oss.link") + "/" + videoDir
	pictureUrl := "http://" + viper.GetString("oss.link") + "/" + pictureDir

	// 等待上传和创建数组库完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 上传视频，切取封面，上传图片
	go func() {
		defer wg.Done()
		// 上传视频
		updataErr = oss7.UploadFileWithByte(videoDir, req.Data)
		// 获取封面,获取第1.0秒的封面
		coverByte, _ := cut.Cover(videoUrl, "00:00:01")
		// 上传封面
		updataErr = oss7.UploadFileWithByte(pictureDir, coverByte)
		logger.Log.Info("上传成功")
	}()

	var videoID int64
	// 创建数据
	go func() {
		defer wg.Done()
		// 创建video
		// CreatedAt ? 不让写，会报错
		video := model.Videos{
			AuthID:   req.UserId,
			Title:    title,
			CoverUrl: pictureUrl,
			PlayUrl:  videoUrl,
		}
		videoID, creatErr = model.GetVideoInstance().Create(video)
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
				// 根据url查找，效率比较低
				// 使用id查找
				_ = model.GetVideoInstance().DeleteVideoByID(videoID)
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

// PublishList 发布列表
func (*VideoService) PublishList(ctx context.Context, req *video_server.PublishListRequest) (resp *video_server.PublishListResponse, err error) {
	resp = new(video_server.PublishListResponse)
	var videos []model.Videos
	key := fmt.Sprintf("%s:%s:%s", "user", "work_list", strconv.FormatInt(req.UserId, 10))

	// 根据用户id找到所有的视频,先找缓存，再查数据库
	exists, err := cache.Redis.Exists(cache.Ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("缓存错误：%v", err)
	}

	if exists > 0 {
		videosString, err := cache.Redis.Get(cache.Ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
		err = json.Unmarshal([]byte(videosString), &videos)
		if err != nil {
			return nil, err
		}
	} else {
		videos, err = model.GetVideoInstance().GetVideoListByUser(req.UserId)
		if err != nil {
			resp.StatusCode = exception.VideoUnExist
			resp.StatusMsg = exception.GetMsg(exception.VideoUnExist)
			return resp, err
		}
		// 放入缓存中
		videosJson, _ := json.Marshal(videos)
		err := cache.Redis.Set(cache.Ctx, key, videosJson, 12*time.Hour).Err()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	resp.VideoList = BuildVideo(videos, req.UserId)

	return resp, nil
}

func BuildVideo(videos []model.Videos, userId int64) []*video_server.Video {
	var videoResp []*video_server.Video

	for _, video := range videos {
		// 查询是否有喜欢的缓存，如果有，比对缓存，如果没有，构建缓存再查缓存
		favorite := isFavorite(userId, int64(video.ID))
		favoriteCount := getFavoriteCount(int64(video.ID))
		commentCount := GetCommentCount(int64(video.ID))
		videoResp = append(videoResp, &video_server.Video{
			Id:            int64(video.ID),
			AuthId:        video.AuthID,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    favorite,
			Title:         video.Title,
		})
	}

	return videoResp
}

func BuildVideoForFavorite(videos []model.Videos, isFavorite bool) []*video_server.Video {
	var videoResp []*video_server.Video

	for _, video := range videos {
		favoriteCount := getFavoriteCount(int64(video.ID))
		commentCount := GetCommentCount(int64(video.ID))
		videoResp = append(videoResp, &video_server.Video{
			Id:            int64(video.ID),
			AuthId:        video.AuthID,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoResp
}

// CountInfo 计数信息
func (*VideoService) CountInfo(ctx context.Context, req *video_server.CountRequest) (resp *video_server.CountResponse, err error) {
	resp = new(video_server.CountResponse)

	userIds := req.UserIds

	for _, userId := range userIds {
		var count video_server.Count

		// 获取赞的数量
		var videos []model.Videos
		exist, err := cache.Redis.HExists(cache.Ctx, "user:total_favorite", strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, fmt.Errorf("1缓存错误：%v", err)
		}

		if exist == false {
			// 获取所有作品数量
			var totalFavorite int64
			totalFavorite = 0
			videos, err = model.GetVideoInstance().GetVideoListByUser(userId)

			for _, video := range videos {
				videoId := video.ID

				favoriteCount, err := model.GetFavoriteInstance().GetVideoFavoriteCount(int64(videoId))
				if err != nil {
					resp.StatusCode = exception.UserNoFavorite
					resp.StatusMsg = exception.GetMsg(exception.UserNoFavorite)
					return resp, err
				}
				log.Print(favoriteCount)
				totalFavorite = totalFavorite + favoriteCount
				log.Print(totalFavorite)
			}
			// 放入缓存
			err = cache.Redis.HSet(cache.Ctx, "user:total_favorite", strconv.FormatInt(userId, 10), totalFavorite).Err()
			if err != nil {
				return nil, err
			}
			cache.Redis.Expire(cache.Ctx, "user:total_favorite", 5*time.Minute)
		} else {
			// 存在缓存
			count.TotalFavorited, err = cache.Redis.HGet(cache.Ctx, "user:total_favorite", strconv.FormatInt(userId, 10)).Int64()
			if err != nil {
				return nil, fmt.Errorf("2缓存错误：%v", err)
			}
		}

		// 获取作品数量
		exist, err = cache.Redis.HExists(cache.Ctx, "user:work_count", strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, fmt.Errorf("3缓存错误：%v", err)
		}
		// 如果存在则读缓存
		if exist {
			count.WorkCount, err = cache.Redis.HGet(cache.Ctx, "user:work_count", strconv.FormatInt(userId, 10)).Int64()
			if err != nil {
				return nil, fmt.Errorf("4缓存错误：%v", err)
			}
		} else {
			// 不存在则查数据库
			count.WorkCount, err = model.GetVideoInstance().GetWorkCount(userId)
			if err != nil {
				resp.StatusCode = exception.UserNoVideo
				resp.StatusMsg = exception.GetMsg(exception.UserNoVideo)
				return resp, err
			}
			// 放入缓存
			err := cache.Redis.HSet(cache.Ctx, "user:work_count", strconv.FormatInt(userId, 10), count.WorkCount).Err()
			if err != nil {
				return nil, fmt.Errorf("5缓存错误：%v", err)
			}
		}

		// 获取喜欢数量
		exist, err = cache.Redis.HExists(cache.Ctx, "user:favorite_count", strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, fmt.Errorf("6缓存错误：%v", err)
		}
		if exist {
			count.FavoriteCount, err = cache.Redis.HGet(cache.Ctx, "user:favorite_count", strconv.FormatInt(userId, 10)).Int64()
			if err != nil {
				return nil, fmt.Errorf("7缓存错误：%v", err)
			}
		} else {
			count.FavoriteCount, err = model.GetFavoriteInstance().GetFavoriteCount(userId)
			if err != nil {
				resp.StatusCode = exception.UserNoFavorite
				resp.StatusMsg = exception.GetMsg(exception.UserNoFavorite)
				return resp, err
			}

			// 放入缓存
			err := cache.Redis.HSet(cache.Ctx, "user:favorite_count", strconv.FormatInt(userId, 10), count.FavoriteCount).Err()
			if err != nil {
				return nil, fmt.Errorf("8缓存错误：%v", err)
			}
		}

		resp.Counts = append(resp.Counts, &count)
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}
