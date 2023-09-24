package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

// TestBetch 获量设置与批量获取
func TestBetch(t *testing.T) {
	var err error

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	err = rdb.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}

	result, _ := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	fmt.Println(result)
}

// TestIncr自增
func TestIncr(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Incr(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值：", val)

	valBy, err := rdb.IncrBy(ctx, "key", 2).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("最新值", valBy)
}

// TestDecr自减
func TestDecr(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Decr(ctx, "key").Result()
	fmt.Println("最新值", val)

	valBy, err := rdb.DecrBy(ctx, "key", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最值值", valBy)
}

// TestDel 批量删除
func TestDel(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	err = rdb.Del(ctx, "key").Err()
	if err != nil {
		panic(err)
	}

	err = rdb.Del(ctx, "key1", "key2", "key3").Err()
	if err != nil {
		panic(err)
	}

}
