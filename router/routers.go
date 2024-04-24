package router

import (
	"restApi/config"
	"restApi/controllers"
	"restApi/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 只允许这个源的请求
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 设置允许的方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 设置允许的请求头部字段
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,    // 允许发送凭证（cookies, HTTP认证及客户端SSL证明等）
		MaxAge:           1728000, // 预检请求的有效期，单位为秒
	}))
	// r.Use(cors.Default())

	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	comments := r.Group("/comments")
	{
		comments.POST("", controllers.CommentsController{}.CreateComment)
	}

	return r

}
