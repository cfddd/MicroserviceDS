package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"utils/snowFlake"
)

type Favorites struct {
	Common
	UserId  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`
	VideoId uint64 `gorm:"column:video_id;type:bigint(20) unsigned;NOT NULL" json:"video_id"`
}

type FavoriteModel struct {
}

var favoriteModel *FavoriteModel
var favoriteOnce sync.Once

func GetFavoriteInstance() *FavoriteModel {
	favoriteOnce.Do(
		func() {
			favoriteModel = &FavoriteModel{}
		})
	return favoriteModel
}

// AddFavorite 创建点赞
func (*FavoriteModel) AddFavorite(tx *gorm.DB, favorite *Favorites) error {
	result := tx.Where("user_id=? AND video_id=?", favorite.UserId, favorite.VideoId).First(&favorite)
	// 发生除没找到记录的其它错误
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// 如果找到了记录，更新is_favorite置为0
	if result.RowsAffected > 0 {
		result = tx.Delete(result)
		if result.Error != nil {
			return result.Error
		}
	} else {
		flake, _ := snowFlake.NewSnowFlake(7, 2)
		favorite.Common.ID = flake.NextId()
		result = tx.Create(&favorite)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

// DeleteFavorite 删除点赞
func (*FavoriteModel) DeleteFavorite(tx *gorm.DB, favorite *Favorites) error {
	result := tx.Where("user_id=? AND video_id=?", favorite.UserId, favorite.VideoId).First(&favorite)
	// 发生除没找到记录的其它错误
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	// 如果找到了记录，更新is_favorite置为0
	if result.RowsAffected > 0 {
		result = tx.Delete(result)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

// FavoriteVideoList 根据用户Id获取所有喜欢的视频id
func (*FavoriteModel) FavoriteVideoList(userId int64) ([]int64, error) {
	var videoIds []int64

	result := DB.Table("favorite").
		Where("user_id = ? AND is_favorite = ?", userId, true).
		Pluck("video_id", &videoIds)
	if result.Error != nil {
		return nil, result.Error
	}

	return videoIds, nil
}

// GetFavoriteCount 获取喜欢数量
func (*FavoriteModel) GetFavoriteCount(userId int64) (int64, error) {
	var count int64

	DB.Table("favorite").
		Where("user_id=? AND is_favorite=?", userId, true).
		Count(&count)

	return count, nil
}

// GetVideoFavoriteCount 获取视频货站数量
func (*FavoriteModel) GetVideoFavoriteCount(videoId int64) (int64, error) {
	var count int64

	DB.Table("favorites").
		Where("video_id=? AND is_favorite=?", videoId, true).
		Count(&count)

	return count, nil
}

// FavoriteUserList 根据视频找到所有点赞用户的id
func (*FavoriteModel) FavoriteUserList(videoId int64) ([]int64, error) {
	var userIds []int64

	result := DB.Table("favorites").
		Where("video_id = ? AND deleted_at is NOT NULL", videoId).
		Pluck("user_id", &userIds)
	if result.Error != nil {
		return nil, result.Error
	}

	return userIds, nil
}
