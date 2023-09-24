package list

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestLPush(t *testing.T) {
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
	//err = rdb.LPush(ctx, "key", "data1").Err()
	err = rdb.LPush(ctx, "key", 1, 2, 3, 4, 5).Err()
	if err != nil {
		panic(err)
	}

	t.Log("push 成功")
}

// TestLrange
func TestLrange(t *testing.T) {
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
	result, err1 := rdb.LRange(ctx, "key", 0, -1).Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// TestLRem删除数据
func TestLRem(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err = rdb.Ping(ctx).Result()
	t.Log("数据库连接成功")
	if err != nil {
		panic(err)
	}
	err = rdb.LRem(ctx, "key", 10, 5).Err()
	if err != nil {
		panic(err)
	}
	t.Log("删除成功")
}

// TestLIndex 获取索引
func TestLIndex(t *testing.T) {
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

	val, err := rdb.LIndex(ctx, "key", 1).Result()
	if err != nil {
		panic(err)
	}
	t.Log(val)
}
