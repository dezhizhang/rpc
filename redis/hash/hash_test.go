package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestHSet(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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
	t.Log("设置值成功")
}

func TestHGet(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log("连接数据库成功")
	result, err1 := rdb.HGet(ctx, "user", "name").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// TestHGetAll 获取所有哈希值
func TestHGetAll(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Logf("连接数据库成功")
	result, err1 := rdb.HGetAll(ctx, "user").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// 哈希累加
func TestHIncrBy(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log("数据库连接成功")
	count, err1 := rdb.HIncrBy(ctx, "user", "count", 2).Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(count)
}

// 获取所有TestHKeys的key

func TestHKeys(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log("连接数据库成功")
	keys, err1 := rdb.HKeys(ctx, "user").Result()
	if err1 != nil {
		panic(err)
	}
	t.Log(keys)

}

// TestHLen 获取hash的长度

func TestLen(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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

// 批量获取
func TestMGet(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	result, err1 := rdb.HMGet(ctx, "user", "name", "count").Result()
	if err1 != nil {
		panic(err)
	}
	t.Log(result)
}

// 批量设置
func TestHMSet(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["name"] = "刘德华"
	data["age"] = 44
	data["gender"] = "男"

	err = rdb.HMSet(ctx, "user", data).Err()
	if err != nil {
		panic(err)
	}
}

func TestHDel(t *testing.T) {
	var err error
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	err = rdb.HDel(ctx, "user", "name").Err()
	if err != nil {
		panic(err)
	}

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
