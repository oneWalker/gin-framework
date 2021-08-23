package service

import (
	"fmt"
	"gin-practice/pkg/db/redis"
)

func StringDemo() {
	//String的操作
	err := redis.RedisDB.Set("name", "kun", 0).Err()
	var val string
	val, _ = redis.RedisDB.Get("name").Result()
	fmt.Println(val)
	//如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则j进行刷新
	redis.RedisDB.SetNX("name2", "lisi", 0).Err()

	//MSET
	redis.RedisDB.MSet("key1", "value1", "key2", "value2", "key3", "value3").Err()
	//MGET
	vals, _ := redis.RedisDB.MGet("key1", "key2", "key3").Result()
	fmt.Println(vals)
	//INCR&DECR
	err = redis.RedisDB.Set("age", "20", 0).Err()
	redis.RedisDB.Incr("age")
	redis.RedisDB.IncrBy("age", 5)
	redis.RedisDB.Decr("age")
	redis.RedisDB.DecrBy("age", 5)
	//Del&Expire&Append
	redis.RedisDB.Append("age", "newAge")
	//seconds
	redis.RedisDB.Expire("age", 5)
	//批量删除
	redis.RedisDB.Del("key1", "key2")
	fmt.Println(err)
}

func ListDemo() {

}

func hashDemo() {

}

func sortSetDemo() {

}

func transaction() {

}
