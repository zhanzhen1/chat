package utils

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Red *redis.Client

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 30,
	})
	result, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("redis init  ", err)
	} else {
		fmt.Println("redis init ", result)
	}
}

const (
	PublishKey = "websocket"
)

func Publish(channel string, msg interface{}) error {
	var err error
	fmt.Println("Publish ", msg)
	err = Red.Publish(channel, msg).Err()
	if err != nil {
		fmt.Println("Publish...", err)
	}
	return err
}
func Subscribe(channel string) (string, error) {
	subscribe := Red.Subscribe(channel)
	fmt.Println("Subscribe...", subscribe)
	message, err := subscribe.ReceiveMessage()
	if err != nil {
		fmt.Println("ReceiveMessage...", err)
	}
	fmt.Println("subscribe ...", message)
	return message.Payload, err
}
