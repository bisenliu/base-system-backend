package initialize

import "C"
import (
	_ "base-system-backend/docs" // 千万不要忘了导入把你上一步生成的docs
	"base-system-backend/global"
	"base-system-backend/middleware"
	"base-system-backend/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	switch global.ENV {
	case "product":
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	case
		"test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	baseRouterGroup := r.Group("v1")
	// 认证中间件
	baseRouterGroup.Use(middleware.GinLogger(), middleware.GinRecovery(true), middleware.JWTAuthMiddleware())
	// 用户模块路由组
	{
		userRouter := router.RouterGroupApp.User
		userRouterGroup := baseRouterGroup.Group("/user/")
		userRouter.InitUserRouter(userRouterGroup)
	}
	// 角色模块路由组
	{
		roleRouter := router.RouterGroupApp.Role
		roleRouterGroup := baseRouterGroup.Group("/role/")
		roleRouter.InitRoleRouter(roleRouterGroup)
	}
	// 日志模块路由组
	{
		logRouter := router.RouterGroupApp.Log
		logRouterGroup := baseRouterGroup.Group("/log/")
		logRouter.InitLogRouter(logRouterGroup)
	}
	// 权限模块路由组
	{
		privilegeRouter := router.RouterGroupApp.Privilege
		privilegeRouterGroup := baseRouterGroup.Group("/privilege/")
		privilegeRouter.InitPrivilegeRouter(privilegeRouterGroup)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r

}
