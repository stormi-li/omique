package omique

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

func NewClient(opts *redis.Options) *Client {
	return &Client{configManager: omi.NewConfigManager(opts)}
}
