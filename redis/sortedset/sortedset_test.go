package sortedset

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestZAdd(t *testing.T) {
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
	err = rdb.ZAdd(ctx, "sortedset", redis.Z{Score: 2.5, Member: "刘德华"}).Err()
	if err != nil {
		panic(err)
	}

}
