// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// 1.创建路由
// 	r := gin.Default()
// 	// 2.绑定路由规则，执行的函数
// 	// gin.Context，封装了request和response
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello World!")
// 	})
// 	// 3.监听端口，默认在8080
// 	// Run("里面不指定端口号默认为8080")
// 	r.Run(":8000")
// }
package main

import (
	"gin-practice/initialize"
	mongodb "gin-practice/pkg/db/mongodb"
	mysql "gin-practice/pkg/db/mysql"
	redis "gin-practice/pkg/db/redis"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//设置配置环境
	//加载.env文件
	//获取运行环境
	//读取配置文件，解析配置文件（放入相应的数据库中）
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file #%v ", err)
	}
	//根据环境更改gin的模式
	env := os.Getenv("ENV")
	if env == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	//mongodb数据库
	if err := mongodb.Init(); err != nil {
		return
	}
	//mysql数据库
	if err := mysql.Init(); err != nil {
		return
	}
	//mysql数据表进行自动生成
	// var foo modelSql.Foo
	// mysql.DB.Begin().AutoMigrate(&foo)

	//redis数据库
	if err := redis.Init(); err != nil {
		return
	}

	//Router的初始化
	r := initialize.Routers()

	r.Run(":8080")
}
