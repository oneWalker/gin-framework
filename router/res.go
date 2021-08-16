package router

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func InitResRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	ResGroup := Router.Group("res")
	{
		ResGroup.GET("resJson", controller.ResJson)
		ResGroup.GET("resStruct", controller.ResStruct)
		ResGroup.GET("resXML", controller.ResXML)
		ResGroup.GET("resYAML", controller.ResYAML)
		ResGroup.GET("resProtobuf", controller.ResProtobuf)
		ResGroup.GET("resTemplate", controller.ResTemplate)
		ResGroup.GET("resRedirect", controller.ResRedirect)
		ResGroup.GET("resAsync", controller.ResAsync)
	}
	return ResGroup
}
