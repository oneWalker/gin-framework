package service

import (
	"fmt"
	"gin-practice/pkg/db/redis"
	"time"

	rawRedis "github.com/go-redis/redis"
)

func StringDemo() {
	//String的操作
	redis.RedisDB.Set("name", "kun", 0).Err()
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
	redis.RedisDB.Set("age", "20", 0).Err()
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
	//统计开发语言排行榜
	zsetKey := "language_rank"
	languages := []rawRedis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}

	// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	num, _ := redis.RedisDB.ZAdd(zsetKey, languages...).Result()

	// 添加一个元素到集合
	redis.RedisDB.ZAdd(zsetKey, rawRedis.Z{Score: 87, Member: "Vue"}).Err()

	//给元素Vue加上8分，最终vue得分95分
	redis.RedisDB.ZIncrBy(zsetKey, 8, "Vue")
	// 返回从0到-1位置的集合元素， 元素按分数从小到大排序 0到-1代表则返回全部数据
	values, _ := redis.RedisDB.ZRange(zsetKey, 0, -1).Result()

	//返回集合元素的个数
	size, _ := redis.RedisDB.ZCard(zsetKey).Result()
	//统计某个分数段内的元素个数，这里是查询的95<分数<100的元素个数
	count, _ := redis.RedisDB.ZCount(zsetKey, "95", "100").Result()

	// 返回从0到-1位置的集合元素， 元素按分数从小到大排序 0到-1代表则返回全部数据
	//ZRevRange分数是从打到小排序，用法和ZRange一样
	values, _ = redis.RedisDB.ZRange(zsetKey, 0, -1).Result()

	// 初始化查询条件， Offset和Count用于分页
	op := rawRedis.ZRangeBy{
		Min:    "80",  // 最小分数
		Max:    "100", // 最大分数
		Offset: 0,     // 类似sql的limit, 表示开始偏移量
		Count:  5,     // 一次返回多少数据
	}
	//根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	values, _ = redis.RedisDB.ZRangeByScore(zsetKey, op).Result()

	// 查询集合元素Golang的分数
	score, _ := redis.RedisDB.ZScore(zsetKey, "Golang").Result()
	//根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	rk, _ := redis.RedisDB.ZRank(zsetKey, "Java").Result()

	fmt.Print(num, values, size, count, score, rk)

	// 删除集合中的元素Java
	redis.RedisDB.ZRem(zsetKey, "Java")

	// 删除集合中的元素Golang和Vue
	// 支持一次删除多个元素
	redis.RedisDB.ZRem(zsetKey, "Golang", "Vue")

	//根据索引范围删除元素
	//集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
	// 这里相当于删除最低分的几个元素
	redis.RedisDB.ZRemRangeByRank(zsetKey, 0, 5)

	// 位置参数写成负数，代表从高分开始删除。
	// 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
	redis.RedisDB.ZRemRangeByRank(zsetKey, -1, -2)
}

func transaction() {
	//以Pipeline的方式操作事务
	//统计开发语言排行榜
	zsetKey := "language_rank"
	// 开启一个TxPipeline事务
	pipe := redis.RedisDB.TxPipeline()

	// 执行事务操作，可以通过pipe读写redis
	incr := pipe.Incr(zsetKey)
	pipe.Expire(zsetKey, time.Hour)

	// 通过Exec函数提交redis事务
	_, err := pipe.Exec()
	//类似于提交
	fmt.Println(incr.Val(), err)

	//Watch - redis乐观锁支持,通过watch监听一些Key,如果这些key的值没有被其他人改变的话,才可以提交事务

	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *rawRedis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get("key").Result()
		if err != nil && err != rawRedis.Nil {
			return err
		}

		// 这里可以处理业务
		fmt.Println(v)

		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(func(pipe rawRedis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set("key", "new value", 0)
			return nil
		})
		return err
	}

	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
	redis.RedisDB.Watch(fn, "key")
}
