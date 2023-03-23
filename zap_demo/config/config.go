package config

import (
	"encoding/json"
	"os"
)

// Config 整个项目的配置；
type Config struct {
	Mode       string `json:"mode"` // 开发、测试和生产环境（debug、test、release）
	Port       int    `json:"port"` // Web 服务运行端口；
	*LogConfig `json:"log"`
	// *other 项目总配置会包括多个细分功能的配置，如LogConfig
}

// Zap 日志库的配置；
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`     // Lumberjack 日志切片相关；
	MaxAge     int    `json:"max_age"`     // Lumberjack 日志切片相关；
	MaxBackups int    `json:"max_backups"` // Lumberjack 日志切片相关；
}

var Conf = new(Config)

func Init(filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf) // 把序列化的b（json字符串或者字符流） 反序列化为结构体实例；
}
