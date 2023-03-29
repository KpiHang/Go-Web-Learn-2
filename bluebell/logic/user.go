package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在  // ParamSignUp是从前端传来参数的结构体；
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	// 生成UID
	userID, _ := snowflake.GetID()
	// 密码加密
	user := &models.User{ // User是将要存放进数据库里的结构体，是从ParamSignUp进行一些操作后来的；
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存到数据库
	return mysql.InsertUser(user)
}
