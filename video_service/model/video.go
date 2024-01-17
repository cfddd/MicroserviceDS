package model

import (
	"gorm.io/gorm"
	"sync"
	"time"
	"utils/snowFlake"
	"video_service/handler"
	video_server "video_service/server"
)

type Common struct {
	ID        uint64         `gorm:"primary_key"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Video struct {
	Common
	AuthID        int64  `gorm:"column:auth_id;type:bigint(20)" json:"auth_id"`
	VideoCreator  int64  `gorm:"column:video_creator;type:bigint(20)" json:"video_creator"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(256)" json:"play_url"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(256)" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:bigint(20)" json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count;type:bigint(20)" json:"comment_count"`
	Title         string `gorm:"column:title;type:varchar(256)" json:"title"`
}

type VideoModel struct {
}

var videoMedel *VideoModel
var videoOnce sync.Once // 单例模式

// GetVideoInstance 获取单例的实例
func GetVideoInstance() *VideoModel {
	videoOnce.Do(
		func() {
			videoMedel = &VideoModel{}
		})
	return videoMedel
}

// GetVideoByTime 根据创建时间获取视频 TODO：后续可以加一点推荐算法
func (*VideoModel) GetVideoByTime(timePoint time.Time) ([]Video, error) {
	var videos []Video

	result := DB.Table("video").
		Where("creat_at < ?", timePoint).
		Order("creat_at DESC").
		Limit(20).
		Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}

	// 查询不到数据，就返回当前时间最新的30条数据
	if len(videos) == 0 {
		timePoint = time.Now()
		result := DB.Table("video").
			Where("creat_at < ?", timePoint).
			Order("creat_at DESC").
			Limit(20).
			Find(&videos)
		if result.Error != nil {
			return nil, result.Error
		}
		return videos, nil
	}

	return videos, nil
}

func BuildVideo(videos []Video, userId int64) []*video_server.Video {
	var videoResp []*video_server.Video

	for _, video := range videos {
		// 查询是否有喜欢的缓存，如果有，比对缓存，如果没有，构建缓存再查缓存
		//favorite := isFavorite(userId, video.Id) //todo：如果useID为-1，那么直接返回false
		//favoriteCount := getFavoriteCount(video.Id)
		commentCount := handler.GetCommentCount(int64(video.ID))
		videoResp = append(videoResp, &video_server.Video{
			Id:       int64(video.ID),
			AuthId:   video.AuthID,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			//FavoriteCount: favoriteCount,
			CommentCount: commentCount,
			//IsFavorite:    favorite,
			Title: video.Title,
		})
	}

	return videoResp
}

// Create 创建视频信息
func (*VideoModel) Create(video *Video) (ID uint64, err error) {
	// 服务2
	flake, _ := snowFlake.NewSnowFlake(7, 2)
	video.ID = uint64(flake.NextId())

	DB.Create(&video)

	return video.ID, nil
}

// DeleteVideoByID 通过ID删除视频
func (v *VideoModel) DeleteVideoByID(id uint64) error {
	var video Video
	if err := DB.Where("id = ?", id).First(&video).Error; err != nil {
		return err
	}

	// 删除找到的记录
	if err := DB.Delete(&video).Error; err != nil {
		return err
	}

	return nil
}

// GetVideoListByUser 根据用户的id找到视频列表
func (*VideoModel) GetVideoListByUser(userId int64) ([]Video, error) {
	var videos []Video

	result := DB.Table("video").
		Where("auth_id = ?", userId).
		Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}

	return videos, nil
}
