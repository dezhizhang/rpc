package publish

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestSubscribe(t *testing.T) {
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
	sub := rdb.Subscribe(ctx, "channel")

	for ch := range sub.Channel() {
		fmt.Println(ch.Channel)
		fmt.Println(ch.Payload)
	}
}

func TestPublish(t *testing.T) {
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
	err = rdb.Publish(ctx, "channel", "hello").Err()
	if err != nil {
		panic(err)
	}
	t.Log("发送消息成历")
}
