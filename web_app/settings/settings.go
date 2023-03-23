package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// use viper to load settings

func Init() (err error) {
	viper.SetConfigName("config")     // specify the config file (without extension);
	viper.SetConfigType("yaml")       // specify the type of the config file;
	viper.AddConfigPath("./settings") // optionally look for config in the working directory;
	err = viper.ReadInConfig()        // Find and read the config file;
	if err != nil {
		// read config file failed;
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	viper.WatchConfig() // 注册监听配置变化的事件；支持热更新；
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
	return
}
