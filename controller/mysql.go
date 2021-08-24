package controller

import (
	model "gin-practice/model/mysql"
	"gin-practice/model/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	foo *model.Foo = &model.Foo{}
)

func AddFoo(ctx *gin.Context) {
	var create *model.Foo
	if err := ctx.BindJSON(&create); err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	id, err := foo.Insert(create)
	if err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	create.ID = uint(id)
	logrus.Info("%v", id)
	response.SUCCESS(create, ctx)
}

func GetFoo(ctx *gin.Context) {
	data, err := foo.GetAll()
	//data, err := foo.GetOne()
	logrus.Info("data", data)
	if err != nil {
		logrus.Fatalf("query faild %v", err)
	}
	response.SUCCESS(data, ctx)
}

func PutFoo(ctx *gin.Context) {
	var put *model.Foo
	if err := ctx.BindJSON(&put); err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	err := foo.Update(put)
	if err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	response.SUCCESS(true, ctx)
}

func DeleteFoo(ctx *gin.Context) {
	ids := ctx.Query("ids")
	id_arr := strings.Split(ids, `,`)
	for _, v := range id_arr {
		num, _ := strconv.Atoi(v)
		foo.Delete(num)
	}
	response.SUCCESS(true, ctx)
}
