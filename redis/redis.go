package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:        "10.10.10.239:6379",
		Password:    "", // 没有密码，默认值
		DB:          0,  // 默认DB 0
		DialTimeout: 5 * time.Second,
	})
	err := rdb.Ping(ctx).Err()
	if err == nil {
		fmt.Println("ping pong")
	} else {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(val)
}
