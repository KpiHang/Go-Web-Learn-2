package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// viper 中使用结构体tag，统一都叫mapstructure;

// Conf 全局变量，用来保存程序的所有配置信息；
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_connections"`
	MaxIdleConns int    `mapstructure:"max_idle_connections"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// use viper to load settings
func Init() (err error) {
	viper.SetConfigName("config") // specify the config file name;
	// viper.SetConfigType("yaml")   // 指定配置文件类型（专用于从远程获取配置信息时指定配置文件类型；）
	viper.AddConfigPath("./conf") // optionally look for config in the working directory;
	err = viper.ReadInConfig()    // Find and read the config file;
	if err != nil {
		// read config file failed;
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到Conf变量中；
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	// 查看配置情况；
	fmt.Printf("%#v\n, %#v\n", Conf, viper.GetInt("app.port"))

	viper.WatchConfig() // 注册监听配置变化的事件；支持热更新；
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("running viper.Unmarshal failed, err:%v\n", err)
			return
		}
		fmt.Println("config file changed:", e.Name)
	})
	return
}
