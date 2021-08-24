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
	var err error
	//仅当列表存在的时候才插入数据,此时列表不存在，无法插入
	redis.RedisDB.LPushX("studentList", "tom")

	//此时列表不存在，依然可以插入
	redis.RedisDB.LPush("studentList", "jack")

	//此时列表存在的时候才能插入数据
	redis.RedisDB.LPushX("studentList", "tom")

	// LPush支持一次插入任意个数据
	err = redis.RedisDB.LPush("studentList", "lily", "lilei", "zhangsan", "lisi").Err()
	if err != nil {
		panic(err)
	}
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, _ := redis.RedisDB.LRange("studentList", 0, -1).Result()
	//返回list集合中的长度
	studentLen, _ := redis.RedisDB.LLen("studentList").Result()

	// 列表索引从0开始计算，这里返回第3个元素
	index, _ := redis.RedisDB.LIndex("studentList", 2).Result()

	//截取名称为key的list,并把截取后的值赋值给studentList
	trim := redis.RedisDB.LTrim("studentList", 0, 3)

	//给名称为key的list中index位置的元素赋值，把原来的数据覆盖
	redis.RedisDB.LSet("studentList", 2, "beer")

	//在list列表studentList中值为lilei前面添加元素hello
	redis.RedisDB.LInsert("studentList", "before", "lilei", "hello")
	//redisDb.LInsertBefore("studentList","lilei","hello") 执行效果同22行

	//在list列表studentList中值为tom后面添加元素world
	redis.RedisDB.LInsert("studentList", "after", "tom", "world")
	//redisDb.LInsertAfter("studentList","tom","world") 执行效果同26行

	//从列表左边删除第一个数据，并返回删除的数据
	redis.RedisDB.LPop("studentList")

	//删除列表中的数据。删除count个key的list中值为value 的元素。如果出现重复元素，仅删除1次，也就是删除第一个
	redis.RedisDB.LRem("studentList", 10, "lisi")
	fmt.Print(vals, studentLen, index, trim) //注意列表是有序的，输出结果是[lisi zhangsan lilei lily tom jack]
}

func hashDemo() {
	var err error
	//SAdd && SCard
	// 添加100,200到集合中
	redis.RedisDB.SAdd("stuSet", 100, 200)
	//获取集合中元素的个数
	size, _ := redis.RedisDB.SCard("stuSet").Result()
	//返回名称为集合中元素的个数
	es, _ := redis.RedisDB.SMembers("stuSet").Result()
	//SisMemeber & SMemebers
	//此处flag1=true
	flag1, _ := redis.RedisDB.SIsMember("stuSet", "200").Result()

	//并集 & 交集 & 差集
	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, _ := redis.RedisDB.SInter("blacklist", "whitelist").Result()
	//求交集并将交集保存到 destSet 的集合
	res, _ := redis.RedisDB.SInterStore("destSet", "blacklist", "whitelist").Result()
	// 求差集
	diffStr, _ := redis.RedisDB.SDiff("blacklist", "whitelist").Result()
	//集合删除操作
	//随机返回集合中的一个元素，并且删除这个元素,这里删除的是400
	member1, _ := redis.RedisDB.SPop("stuSet").Result()
	// 随机返回集合中的4个元素，并且删除这些元素
	member2, _ := redis.RedisDB.SPopN("stuSet", 4).Result()
	//删除集合stuSet名称为300,400的元素,并返回删除的元素个数
	member3, _ := redis.RedisDB.SRem("stuSet", 500, 600).Result()
	//集合随机数操作
	//随机返回集合stuSet中的一个元素
	member4, _ := redis.RedisDB.SRandMember("stuSet").Result()
	//随机返回集合stuSet中的3个元素
	member5, _ := redis.RedisDB.SRandMemberN("stuSet", 3).Result()
	fmt.Print(err, size, es, flag1, names, res, diffStr, member1, member2, member3, member4, member5)
}

func sortSetDemo() {

}

func transaction() {

}
