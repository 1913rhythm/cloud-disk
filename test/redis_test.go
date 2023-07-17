package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "124.221.4.71:6379",
	Password: "123456",
	DB:       0,
})

// 向redis中插入数据
func TestRedisSet(t *testing.T) {
	err := rdb.Set(ctx, "name", "jack", time.Second*30).Err()
	if err != nil {
		t.Error(err)
	}
}

// 从redis中获取数据
func TestRedisGet(t *testing.T) {
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)
}
