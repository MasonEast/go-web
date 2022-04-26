package route

import (
	"myapp/api"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	var admin api.AdminApi
	r.GET("/admin", admin.Get)

	// Listen and Server in 0.0.0.0:8888
	r.Run(":8888")
}