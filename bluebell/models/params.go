package models

// 定义请求的参数结构体；

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// binding 是gin框架中shouldbind 到结构体中会用；
