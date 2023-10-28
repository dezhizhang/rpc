package sortedset

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

// 添加一个或多个元素到集合，如果元素已经存在则更新分数
func TestZAdd(t *testing.T) {

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
	err = rdb.ZAdd(ctx, "zAdd", redis.Z{Score: 2.5, Member: "张德志"}).Err()
	if err != nil {
		panic(err)
	}
	t.Log("插入成功")
}

// 返回集合元素个数
func TestZCard(t *testing.T) {
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
	size, err1 := rdb.ZCard(ctx, "zAdd").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(size)
}

// 获取某个区间的值
func TestZCount(t *testing.T) {
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
	size, err1 := rdb.ZCount(ctx, "zAdd", "(1", "5").Result()
	if err1 != nil {
		panic(err)
	}
	t.Log(size)
}

func TestZIncrBy(t *testing.T) {
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
	err = rdb.ZIncrBy(ctx, "zAdd", 2, "张德志").Err()
	if err != nil {
		panic(err)
	}
}

// 返回集合中某个索引范围的元素
func TestZRange(t *testing.T) {
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
	result, err1 := rdb.ZRange(ctx, "zAdd", 0, -1).Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(result)
}
