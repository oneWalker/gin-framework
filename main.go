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
	"context"
	"gin-practice/initialize"
	mongodb "gin-practice/pkg/db/mongodb"
	"net/http"
	"os/signal"
	"time"

	//mysql "gin-practice/pkg/db/mysql"
	//redis "gin-practice/pkg/db/redis"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	//设置配置环境
	//加载.env文件
	//获取运行环境
	//读取配置文件，解析配置文件（放入相应的数据库中）
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file #%v ", err)
	}
	//根据环境更改gin的模式
	env := os.Getenv("ENV")
	if env == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	//测试mongo相关demo暂时不使用
	//mongodb数据库
	if err := mongodb.Init(); err != nil {
		return
	}
	// //redis数据库
	// if err := redis.Init(); err != nil {
	// 	return
	// }

	// //mysql数据库
	// if err := mysql.Init(); err != nil {
	// 	return
	// }
	//mysql数据表进行自动生成
	// var foo modelSql.Foo
	// mysql.DB.Begin().AutoMigrate(&foo)

	// //Router的初始化，原启动方式
	// r := initialize.Routers()
	// r.Run(":8080")

	//RPC相关的初始化
	//初始化rpc服务
	// srv := grpc.NewServer()

	//初始化grpc服务
	//初始化rpc客户端

	//HTTP的优雅启动和重启方式
	srv := &http.Server{
		Addr:    ":8080",
		Handler: initialize.Routers(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Println("Server exiting")

}
