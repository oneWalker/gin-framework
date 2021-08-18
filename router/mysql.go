package router

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func InitMysqlRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	BaseGroup := Router.Group("mysql")
	{
		BaseGroup.POST("addFoo", controller.AddFoo)
		BaseGroup.GET("getFoo", controller.GetFoo)
		BaseGroup.PUT("putFoo", controller.PutFoo)
		BaseGroup.DELETE("deleteFoo", controller.DeleteFoo)
	}
	return BaseGroup
}
