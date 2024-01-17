package handler

import (
	"fmt"
	"testing"
	"time"
	"video_service/model"
)

func TestCommon(t *testing.T) {
	cfd := model.Video{
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
