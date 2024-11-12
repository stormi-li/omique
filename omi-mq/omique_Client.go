package omique

import (
	manager "github.com/stormi-li/omi-v1/omi-manager"
)

type Client struct {
	configManager *manager.Client
}

func (c *Client) NewConsumer(channel string, weight int) *Consumer {
	return &Consumer{
		configManager: c.configManager,
		channel:       channel,
		weight:        weight,
		messageChan:   make(chan []byte, 1000000),
	}
}

func (c *Client) NewProducer(channel string) *Producer {
	producer := Producer{
		configSearcher: c.configManager.NewSearcher(),
		channel:        channel,
	}
	return &producer
}
