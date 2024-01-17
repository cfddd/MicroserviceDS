package redis

import (
	"context"
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitRedis()
	Redis.Set(context.Background(), "k2", "myValue", 0)
}
