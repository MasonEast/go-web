package route

import (
	"myapp/api"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// 公用路由（不需要鉴权）
	{
		publicGroup := r.Group("")
		// 健康检测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		var user api.PublicApi
		publicGroup.POST("/register", user.Register)
		publicGroup.POST("/login", user.Login)

	}

	// 需要鉴权的路由
	{
		privateGroup := r.Group("")
		privateGroup.Use(middleware.Jwt())

		var admin api.AdminApi
		privateGroup.GET("/admin", admin.Get)

		var article api.ArticleApi
		articleGroup := privateGroup.Group("/article")
		articleGroup.POST("/create", article.Create)
	}

	// Listen and Server in 0.0.0.0:8888
	r.Run(":8888")
}