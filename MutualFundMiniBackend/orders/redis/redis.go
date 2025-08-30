package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedisClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to Redis: %v", err))
	}
	fmt.Println("Redis connected")
}

func SetNAV(schemeCode string, nav float64) error {
	key := fmt.Sprintf("nav:latest:%s", schemeCode)
	return Rdb.Set(Ctx, key, nav, 24*time.Hour).Err()
}

func GetNAV(schemeCode string) (float64, error) {
	key := fmt.Sprintf("nav:latest:%s", schemeCode)
	val, err := Rdb.Get(Ctx, key).Float64()
	if err != nil {
		return 0, err
	}
	return val, nil
}
