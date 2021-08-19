package cache

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var Redis *redis.Client

func Init() error {
	var err error
	db, _ := strconv.ParseUint("0", 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       int(db),
	})

	_, err = client.Ping().Result()

	if err != nil {
		logrus.Fatalf("redis connected failed: %v", err)

	}

	Redis = client
	return err
}

func Close() error {
	if Redis != nil {
		if err := Redis.Close(); err != nil {
			return err
		}
	}
	logrus.Println("mysql connect closed")
	return nil
}
