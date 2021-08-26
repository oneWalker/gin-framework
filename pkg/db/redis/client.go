package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var RedisDB *redis.Client

func Init() error {
	var err error
	db, _ := strconv.ParseUint("0", 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       int(db),
		PoolSize: 10, //默认10个
	})
	_, err = client.Ping().Result()

	if err != nil {
		logrus.Fatalf("redis connected failed: %v", err)

	}

	RedisDB = client
	logrus.Info("redis connect successfully")
	return err
}

func Close() error {
	if RedisDB != nil {
		if err := RedisDB.Close(); err != nil {
			return err
		}
	}
	logrus.Println("redis connect closed")
	return nil
}
