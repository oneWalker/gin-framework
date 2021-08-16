package initialize

import (
	//"gin-practice/middleware"
	"gin-practice/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	//设置相应的跨域方式
	//方式1
	Router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept-Encoding", "Authorization", "X-Request-Id", "x-app-name"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//方式2
	//Router.Use(middleware.CrosHandler())

	//加载HTML模版
	Router.LoadHTMLFiles("templates/*")

	//分组进行权限校验
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
		router.InitResRouter(PublicGroup)
	}

	// PrivateGroup := Router.Group("")
	// PrivateGroup.Use(middleware.JWTAuth())
	// {
	// 	router.InitUserRouter(PrivateGroup)
	// 	router.InitTvRouter(PrivateGroup)
	// }

	// YsxGroup := Router.Group("")
	// YsxGroup.Use(middleware.CheckXAcessToken())
	// {
	// 	router.InitContactRouter(YsxGroup)
	// }

	/*
		YsxTokenGroup := Router.Group("")
		YsxTokenGroup.Use(middleware.CheckXAcessToken())
		{
			router.InitContactTokenRouter(YsxTokenGroup)
		}
	*/

	return Router
}
