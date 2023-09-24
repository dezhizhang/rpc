package set

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestSAdd(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err = rdb.Ping(ctx).Result()
	err = rdb.SAdd(ctx, "set", 100).Err()
	if err != nil {
		panic(err)
	}
	t.Log("添加集合成功")
}

func TestSCard(t *testing.T) {
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
	t.Log("连接数据库成功")
	size, err := rdb.SCard(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	t.Log(size)
}

func TestSMembers(t *testing.T) {
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
	values, err := rdb.SMembers(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	t.Log(values)
}
