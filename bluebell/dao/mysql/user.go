package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

// ChechUserExist 检查用户名是否已经存在
func CheckUserExist(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username = ?`
	var count int // 小写db，只在mysql中使用，安全性；
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密；
	user.Password = encryptPassword(user.Password)
	// 执行sql语句，插入数据；
	sqlStr := `INSERT INTO user(user_id, username, password) VALUES(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

var secret = "bluebell"

// md5是一个单向加密算法，它只能加密单向的数据。
// encryptPassword 函数用于将原始密码 oPassword 进行 MD5 加密，并返回加密后的密码。
func encryptPassword(oPassword string) string {
	h := md5.New()          // 创建一个 MD5 散列算法对象
	h.Write([]byte(secret)) // 添加一个字符串作为盐值到散列器中
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
	// 将密码转换为字节切片后加密，并将结果转换为十六进制编码的字符串
}
