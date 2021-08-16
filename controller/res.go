package controller

import (
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
