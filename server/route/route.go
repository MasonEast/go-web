package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name, action)
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}