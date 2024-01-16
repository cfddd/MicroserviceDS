package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	Redis *redis.Client
	ctx   = context.Background()
)

func GenerateFollowKey(UserId int64) string {
	return "FOLLOW:" + strconv.FormatInt(UserId, 10)
}

func GenerateFollowerKey(UserId int64) string {
	return "FOLLOWER:" + strconv.FormatInt(UserId, 10)
}

func GenerateFriendKey(UserId int64) string {
	return "FRIEND:" + strconv.FormatInt(UserId, 10)
}
