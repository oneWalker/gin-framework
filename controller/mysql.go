package controller

import (
	model "gin-practice/model/mysql"
	"gin-practice/model/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddFoo(ctx *gin.Context) {

}

func GetFoo(ctx *gin.Context) {
	var foo *model.Foo
	data, err := foo.Get()
	if err != nil {
		logrus.Fatalf("query faild %v", err)
	}
	response.SUCCESS(data, ctx)
}

func PutFoo(ctx *gin.Context) {

}

func DeleteFoo(ctx *gin.Context) {

}
