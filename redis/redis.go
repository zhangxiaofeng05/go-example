package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 建立 Redis 连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口
		Password: "",               // Redis 认证密码
		DB:       0,                // Redis 数据库索引
	})

	// 检查连接是否成功
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pingRes, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ping value: %v\n", pingRes)

	// 设置键值对
	err = client.Set(ctx, "rose", "jack", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值对
	val, err := client.Get(ctx, "rose").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rose value: %v\n", val)
}
