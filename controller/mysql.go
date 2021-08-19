package controller

import (
	model "gin-practice/model/mysql"
	"gin-practice/model/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	foo *model.Foo = &model.Foo{}
)

func AddFoo(ctx *gin.Context) {

}

func GetFoo(ctx *gin.Context) {
	data, err := foo.GetAll()
	logrus.Info("data", data)
	if err != nil {
		logrus.Fatalf("query faild %v", err)
	}
	response.SUCCESS(data, ctx)
}

func PutFoo(ctx *gin.Context) {

}

func DeleteFoo(ctx *gin.Context) {

}
