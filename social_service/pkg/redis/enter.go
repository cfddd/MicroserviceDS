package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	Redis *redis.Client
	ctx   = context.Background()
)
