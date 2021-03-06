package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-api/api"
	"go-api/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.NoMethod(api.HandleNotFound)
	r.NoRoute(api.HandleNotFound)

	if gin.Mode() != gin.ReleaseMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 中间件, 顺序不能改
	// cors zaplog time/rate
	r.Use(middleware.Cors(),
		middleware.GinLogger(),
		middleware.Rate(),
	)

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// oss token
		v1.GET("oss", api.GetOssToken)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.JWTAuth())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			//refresh token
			auth.PUT("user/token/refresh", api.UserTokenRefresh)
		}
	}
	return r
}
