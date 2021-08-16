package router

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	BaseGroup := Router.Group("base")
	{
		BaseGroup.GET("getTest", controller.GetTest)
		BaseGroup.POST("postTestForm", controller.PostTestForm)
		BaseGroup.POST("postTestJson", controller.PostTestJson)
		BaseGroup.POST("testUri/:name/*password", controller.TestUri)
	}
	return BaseGroup
}
