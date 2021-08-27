package controller

import (
	"gin-practice/model/mongodb"
	"gin-practice/model/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	id, _ := user.InsertOne(&U)
	response.SUCCESS(id, ctx)
}

func GetUser(ctx *gin.Context) {
	//通过结构体进行获取
	var U mongodb.User
	if err := ctx.Bind(&U); err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	//res, err := user.GetOne("61276519aa45955d83b923ca")
	res, err := user.GetAll()
	if err != nil {
		logrus.Fatalf("query error %v", err)
		return
	}
	response.SUCCESS(res, ctx)
}

func PutUser(ctx *gin.Context) {
	//通过结构体进行获取
	var U mongodb.User
	if err := ctx.Bind(&U); err != nil {
		logrus.Fatalf("parma auth faild %v", err)
		return
	}
	filter := bson.D{primitive.E{Key: "id", Value: ""}}
	updateResult, err := user.Update(filter, &U)
	if err != nil {
		logrus.Fatalf("update error %v", err)
		return
	}
	response.SUCCESS(&updateResult, ctx)
}

func DeleteUser(ctx *gin.Context) {

}
