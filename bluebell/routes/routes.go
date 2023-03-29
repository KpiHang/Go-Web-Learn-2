package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode { // 设值模式，dev、release
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
