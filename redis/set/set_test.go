package set

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

// 添加集合
func TestSAdd(t *testing.T) {
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
	err = rdb.SAdd(ctx, "set", 100, 200, 300, 400, 10, 20).Err()
	if err != nil {
		panic(err)
	}
	t.Log("添加集合成功")
}

// 获取集合元素个数
func TestSCard(t *testing.T) {
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
	size, err1 := rdb.SCard(ctx, "set").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(size)

}

// 判断元素是否在集合是
func TestSIsMember(t *testing.T) {
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
	ok, _ := rdb.SIsMember(ctx, "set", 100).Result()
	if !ok {
		t.Log("集合不含令指定元素")
		return
	}
	t.Log("集合包含指定元素")
}

// 获取集合中所有的元素
func TestSMembers(t *testing.T) {
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
	values, err1 := rdb.SMembers(ctx, "set").Result()
	if err1 != nil {
		panic(err1)
	}
	t.Log(values)
}

// 删除集合中元素
func TestSRem(t *testing.T) {
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
	err = rdb.SRem(ctx, "set", 100).Err()
	t.Log("删除成功")
}

// 随机删除并返回删除的值
func TestSPop(t *testing.T) {
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
	val, _ := rdb.SPop(ctx, "set").Result()
	t.Log(val)
	vals, _ := rdb.SPopN(ctx, "set", 5).Result()
	t.Log(vals)
}
