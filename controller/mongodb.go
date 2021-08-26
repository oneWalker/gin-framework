package controller

import (
	"gin-practice/model/mongodb"
	"gin-practice/model/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	user *mongodb.User = &mongodb.User{}
)

func AddUser(ctx *gin.Context) {
	//通过结构体进行获取
	var U mongodb.User
	if err := ctx.Bind(&U); err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	user.InsertOne(&U)
	response.SUCCESS(&U, ctx)
}

func GetUser(ctx *gin.Context) {

}

func PutUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}
