package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"utils/exception"
	"video_service/model"
	"video_service/pkg/cache"
	video_server "video_service/server"
)

// FavoriteAction 点赞操作 todo 也可以设计成定时任务
func (*VideoService) FavoriteAction(ctx context.Context, req *video_server.FavoriteActionRequest) (resp *video_server.FavoriteActionResponse, err error) {
	resp = new(video_server.FavoriteActionResponse)
	key := fmt.Sprintf("%s:%s", "user", "favorite_count")
	setKey := fmt.Sprintf("%s:%s:%s", "video", "favorite_video", strconv.FormatInt(req.VideoId, 10))
	favoriteKey := fmt.Sprintf("%s:%s:%s", "user", "favorit_list", strconv.FormatInt(req.UserId, 10))

	action := req.ActionType
	var favorite model.Favorites
	favorite.UserId = uint64(req.UserId)
	favorite.VideoId = uint64(req.VideoId)

	// 查看缓存是否存在，不存在这构建一次缓存，避免极端情况
	setExists, err := cache.Redis.Exists(cache.Ctx, setKey).Result()
	if err != nil {
		return nil, fmt.Errorf("缓存错误：%v", err)
	}
	if setExists == 0 {
		// 构建点赞这个视频的用户集合的缓存，使用redis的set集合存储缓存
		err := buildVideoFavorite(req.VideoId)
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
	}

	// 点赞操作
	if action == 1 {
		// 查看缓存的set集合中是否有这个用户的点赞记录，避免重复点赞
		result, err := cache.Redis.SIsMember(cache.Ctx, setKey, req.UserId).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		if result {
			// 重复点赞
			resp.StatusCode = exception.FavoriteErr
			resp.StatusMsg = exception.GetMsg(exception.FavoriteErr)
			return resp, err
		}

		// 操作favorite（点赞）表
		// 此处使用事务，避免出现数据库操作成功，但是缓存没有增加成功的情况？
		tx := model.DB.Begin()
		err = model.GetFavoriteInstance().AddFavorite(tx, &favorite)
		if err != nil {
			//tx.Rollback()
			resp.StatusCode = exception.FavoriteErr
			resp.StatusMsg = exception.GetMsg(exception.FavoriteErr)
			return resp, err
		}

		// 点赞成功，缓存中点赞总数 + 1
		// 检查当前用户是否有点赞记录
		exist, err := cache.Redis.HExists(cache.Ctx, key, strconv.FormatInt(req.UserId, 10)).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		if exist {
			// 字段存在，该记录数量 + 1
			// 如果当前用户有点赞记录，就将当前用户的点赞数缓存（哈希实现的）进行+1操作
			_, err = cache.Redis.HIncrBy(cache.Ctx, key, strconv.FormatInt(req.UserId, 10), 1).Result()
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("缓存错误：%v", err)
			}
		}

		// 加入喜欢set中，如果没有，构建缓存再加入set中
		// 并且将当前点赞的视频加入的视频喜欢集合的缓存中
		err = cache.Redis.SAdd(cache.Ctx, setKey, strconv.FormatInt(req.UserId, 10)).Err()
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		// 删除喜欢列表缓存
		err = cache.Redis.Del(cache.Ctx, favoriteKey).Err()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
		defer func() {
			go func() { // 另起协程
				//演示双删，延时3秒执行
				time.Sleep(time.Second * 3)
				//再次删除缓存
				cache.Redis.Del(cache.Ctx, favoriteKey)
			}()
		}()
		// 提交事务
		tx.Commit()
	}

	// 取消赞操作
	if action == 2 {
		// 查看缓存，避免重复点删除
		result, err := cache.Redis.SIsMember(cache.Ctx, setKey, req.UserId).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		if result == false {
			// 重复删除
			resp.StatusCode = exception.CancelFavoriteErr
			resp.StatusMsg = exception.GetMsg(exception.CancelFavoriteErr)
			return resp, err
		}

		// 操作favorite表
		tx := model.DB.Begin()
		err = model.GetFavoriteInstance().DeleteFavorite(tx, &favorite)
		if err != nil {
			resp.StatusCode = exception.CancelFavoriteErr
			resp.StatusMsg = exception.GetMsg(exception.CancelFavoriteErr)
			return resp, err
		}

		// 点赞成功，缓存中点赞总数 - 1
		exist, err := cache.Redis.HExists(cache.Ctx, key, strconv.FormatInt(req.UserId, 10)).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		if exist {
			// 字段存在，该记录数量 + 1
			_, err = cache.Redis.HIncrBy(cache.Ctx, key, strconv.FormatInt(req.UserId, 10), -1).Result()
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("缓存错误：%v", err)
			}
		}

		// set中删除
		err = cache.Redis.SRem(cache.Ctx, setKey, req.UserId).Err()
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("缓存错误：%v", err)
		}

		// 删除喜欢列表缓存
		err = cache.Redis.Del(cache.Ctx, favoriteKey).Err()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
		defer func() {
			go func() {
				//延时3秒执行
				time.Sleep(time.Second * 3)
				//再次删除缓存
				cache.Redis.Del(cache.Ctx, favoriteKey)
			}()
		}()

		tx.Commit()
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)

	return resp, nil
}

// FavoriteList 喜欢列表
func (*VideoService) FavoriteList(ctx context.Context, req *video_server.FavoriteListRequest) (resp *video_server.FavoriteListResponse, err error) {
	resp = new(video_server.FavoriteListResponse)
	var videos []model.Videos
	key := fmt.Sprintf("%s:%s:%s", "user", "favorit_list", strconv.FormatInt(req.UserId, 10))

	exits, err := cache.Redis.Exists(cache.Ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("缓存错误：%v", err)
	}

	if exits > 0 {
		videosString, err := cache.Redis.Get(cache.Ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
		err = json.Unmarshal([]byte(videosString), &videos)
		if err != nil {
			return nil, err
		}
	} else {
		// 根据用户id找到所有的视频id
		var videoIds []int64
		videoIds, err = model.GetFavoriteInstance().FavoriteVideoList(req.UserId)
		if err != nil {
			resp.StatusCode = exception.UserNoVideo
			resp.StatusMsg = exception.GetMsg(exception.UserNoVideo)
			return resp, err
		}

		// 根据视频id找到视频的详细信息
		videos, err = model.GetVideoInstance().GetVideoList(videoIds)
		if err != nil {
			resp.StatusCode = exception.VideoUnExist
			resp.StatusMsg = exception.GetMsg(exception.VideoUnExist)
			return resp, err
		}

		// 放入缓存中
		videosJson, _ := json.Marshal(videos)
		err := cache.Redis.Set(cache.Ctx, key, videosJson, 30*time.Minute).Err()
		if err != nil {
			return nil, fmt.Errorf("缓存错误：%v", err)
		}
	}

	resp.StatusCode = exception.SUCCESS
	resp.StatusMsg = exception.GetMsg(exception.SUCCESS)
	resp.VideoList = BuildVideoForFavorite(videos, true)

	return resp, nil
}

// 查询缓存，判断是否喜欢
func isFavorite(userId int64, videoId int64) bool {
	var isFavorite bool
	key := fmt.Sprintf("%s:%s:%s", "video", "favorite_video", strconv.FormatInt(videoId, 10))

	exists, err := cache.Redis.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Print(err)
	}

	if exists > 0 {
		isFavorite, err = cache.Redis.SIsMember(cache.Ctx, key, strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			log.Print(err)
		}
	} else {
		err := buildVideoFavorite(videoId)
		if err != nil {
			log.Print(err)
		}
		isFavorite, err = cache.Redis.SIsMember(cache.Ctx, key, strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			log.Print(err)
		}
	}

	return isFavorite
}

// 构建视频点赞缓存
func buildVideoFavorite(videoId int64) error {
	key := fmt.Sprintf("%s:%s:%s", "video", "favorite_video", strconv.FormatInt(videoId, 10))

	// 查询出所有喜欢这个视频的所有用户id
	userIdList, err := model.GetFavoriteInstance().FavoriteUserList(videoId)
	if err != nil {
		return err
	}
	// 如果点赞数量为空，则不走缓存
	if len(userIdList) == 0 {
		return nil
	}

	userIds := make([]interface{}, len(userIdList))
	for i, video := range userIdList {
		userIds[i] = video
	}

	err = cache.Redis.SAdd(cache.Ctx, key, userIds...).Err()
	if err != nil {
		return err
	}

	return nil
}

// 通过缓存查询视频的获赞数量
func getFavoriteCount(videoId int64) int64 {
	setKey := fmt.Sprintf("%s:%s:%s", "video", "favorite_video", strconv.FormatInt(videoId, 10))

	// 查看缓存是否存在，不存在这构建一次缓存，避免极端情况
	setExists, err := cache.Redis.Exists(cache.Ctx, setKey).Result()
	if err != nil {
		log.Print(err)
	}

	if setExists == 0 {
		err := buildVideoFavorite(videoId)
		if err != nil {
			log.Print(err)
		}
	}

	count, err := cache.Redis.SCard(cache.Ctx, setKey).Result()
	if err != nil {
		log.Print(err)
	}
	return count
}
