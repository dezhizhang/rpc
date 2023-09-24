package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	var err error
	var result interface{}
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "name", "哈哈大", 0).Err()
	if err != nil {
		panic(err)
	}

	result, err = rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}
