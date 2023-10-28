package client

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestSetKey(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Set(ctx, "name", "hello world", time.Second*1000).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("设置值成功")
}

func TestGetKey(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	result, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("获取值", result)
}

func TestGetSet(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	result, err := rdb.GetSet(ctx, "name", "hello change info").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("获取到的值", result)
}

func TestSetNx(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.SetNX(ctx, "name", "set nx", time.Second*1000).Err()
	if err != nil {
		panic(err)
	}
}

// 批量获取
func TestMGet(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	result, err := rdb.MGet(ctx, "name", "k1", "k2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestMSet(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.MSet(ctx, "k1", "value1", "k2", "value2", "k3", "value3").Err()
	if err != nil {
		panic(err)
	}
}

// 自增
func TestIncrBy(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	result, err := rdb.IncrBy(ctx, "money", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// 自减

func TestDecrBy(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	val, err := rdb.DecrBy(ctx, "money", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

// 删除
func TestDel(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Del(ctx, "k1").Err()
	if err != nil {
		panic(err)
	}
}

func TestExpire(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdb.Expire(ctx, "key2", 1000*time.Second)
}
