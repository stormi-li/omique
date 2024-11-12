package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	omique "github.com/stormi-li/omique"
)

func main() {
	producer()
}

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func producer() {
	c := omique.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	producer := c.NewProducer("consumer_1")
	now := time.Now()
	for i := 0; i < 100000; i++ {
		producer.Publish([]byte("message" + strconv.Itoa(i)))
		time.Sleep(10*time.Millisecond)
	}
	fmt.Println(time.Since(now))
}
