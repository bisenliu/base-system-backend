package core

import (
	"base-system-backend/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"path"
)

// Viper
//
//	@Description: 加载配置文件
//	@return *viper.Viper viper对象
func Viper() *viper.Viper {
	v := viper.New()
	currentPath, _ := os.Getwd()
	configPath := path.Join(currentPath, "config.yaml")
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf(" v.ReadInConfig() failed: %w \n", err))
	}

	if err := mapstructure.Decode(v.Get(global.ENV), &global.CONFIG); err != nil {
		panic(fmt.Errorf("mapstructure.Decode failed: %w \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := mapstructure.Decode(v.Get(global.ENV), &global.CONFIG); err != nil {
			panic(fmt.Errorf("mapstructure.Decode failed: %w \n", err))
		}
	})
	return v
}
