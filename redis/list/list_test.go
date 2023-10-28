package list

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

// 左侧插入数据
func TestLPush(t *testing.T) {
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
	err = rdb.LPush(ctx, "key", 1, 2, 3, 4, 5).Err()
	if err != nil {
		panic(err)
	}
	t.Log("插入成功")
}

// 判断集合左侧是否可以插入
func TestLPushX(t *testing.T) {
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
	err = rdb.LPushX(ctx, "key", 6, 7, 8).Err()
	if err != nil {
		panic(err)
	}
}

// 从右则删除一个值并返回值
func TestRPop(t *testing.T) {
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
	val, err1 := rdb.RPop(ctx, "key").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(val)
}

// 从列表右则插入值
func TestRPush(t *testing.T) {
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
	err = rdb.RPush(ctx, "key", 12).Err()
	if err != nil {
		panic(err)
	}
}

// 从左则删除
func TestLPop(t *testing.T) {
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
	result, err1 := rdb.LPop(ctx, "key").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// 获取集合的长度
func TestLLen(t *testing.T) {
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
	result, err1 := rdb.LLen(ctx, "key").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}

// 遍历list
func TestLRange(t *testing.T) {
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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log("数据库连接成功")
	err = rdb.LRem(ctx, "key", 0, -1).Err()
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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	val, err1 := rdb.LIndex(ctx, "key", 1).Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(val)
}
