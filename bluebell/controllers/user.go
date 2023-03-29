package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应；
		zap.L().Error("SignUpHandler failed", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型的；
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error_zh": errs.Translate(trans), // 对错误进行翻译；
		})
		return
	}
	// 2. 业务处理：参数和请求转发给logic层；
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3. 返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
