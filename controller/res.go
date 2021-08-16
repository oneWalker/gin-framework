package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func ResJson(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "someJson", "status": 200})
}

func ResStruct(ctx *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Memner  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Memner = 123
	ctx.JSON(200, msg)
}

func ResXML(ctx *gin.Context) {
	ctx.XML(200, gin.H{"message": "abc"})
}

func ResYAML(ctx *gin.Context) {
	ctx.YAML(200, gin.H{"name": "zhangsan"})
}

func ResProtobuf(ctx *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数据
	label := "label"
	// 传protobuf数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	ctx.ProtoBuf(200, data)
}

func ResTemplate(ctx *gin.Context) {
	ctx.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
}

func ResRedirect(ctx *gin.Context) {
	// 支持内部和外部重定向
	ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}

//异步：正常情况下就是同步的
func ResAsync(ctx *gin.Context) {
	// 需要一个副本
	copyContext := ctx.Copy()
	// 实现异步处理
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("异步执行：" + copyContext.Request.URL.Path)
	}()
}
