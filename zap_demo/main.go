package main

import (
	"fmt"
	"net/http"
	"os"
	"zap_demo/config"
	"zap_demo/logger"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config from config.json // 不指定时加载默认配置，在之后的viper库中实现；
	if len(os.Args) < 1 {
		return
	}
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	gin.SetMode(config.Conf.Mode) // 开发、测试和生产环境（debug、test、release）

	r := gin.Default()
	// 注册zap相关路由；
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/hello", func(c *gin.Context) {
		// 假设你有一些数据需要记录到日志中
		var (
			name = "q1mi"
			age  = 18
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello liwenzhou.com")
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	r.Run(addr)
}
