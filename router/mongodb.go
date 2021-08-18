package router

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func InitMongodbRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	BaseGroup := Router.Group("mongo")
	{
		BaseGroup.POST("addUser", controller.AddUser)
		BaseGroup.GET("getUser", controller.GetUser)
		BaseGroup.PUT("putUser", controller.PutUser)
		BaseGroup.DELETE("deleteUser", controller.DeleteUser)
	}
	return BaseGroup
}
