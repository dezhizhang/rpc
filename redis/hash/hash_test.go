package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

// TestHset 设置hash
func TestHset(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log(result)

	err = rdb.HSet(ctx, "user", "name", "张德志").Err()
	if err != nil {
		panic(err)
	}

	t.Logf("设置值成功")

}

// TestHGet 获取哈希
func TestHGet(t *testing.T) {
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
	t.Logf("连接数据库成功")

	result, err1 := rdb.HGet(ctx, "user", "name").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Logf(result)

}

// TestHGetAll 获取所有哈希值
func TestHGetAll(t *testing.T) {
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
	result, err1 := rdb.HGetAll(ctx, "user").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// TestHIncrBy 哈希累加
func TestHIncrBy(t *testing.T) {
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
	count, err1 := rdb.HIncrBy(ctx, "user", "count", 2).Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(count)
}

// 获取所有TestHkeys的key
func TestHkeys(t *testing.T) {
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

	keys, err1 := rdb.HKeys(ctx, "user").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(keys)
}

// TestHLen 获取hash的长度
func TestHLen(t *testing.T) {
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
	result, err1 := rdb.HLen(ctx, "user").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)

}

// TestHExists 判断hash值是否存在
func TestHExists(t *testing.T) {
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
	result, err1 := rdb.HExists(ctx, "user", "name").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}
