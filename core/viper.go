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

func Viper() *viper.Viper {
	v := viper.New()
	currentPath, _ := os.Getwd()
	configPath := path.Join(currentPath, "base-system-backend-config.yaml")
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf(" v.ReadInConfig() failed: %s \n", err))
	}

	if err := mapstructure.Decode(v.Get(global.ENV), &global.CONFIG); err != nil {
		panic(fmt.Errorf("mapstructure.Decode failed: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := mapstructure.Decode(v.Get(global.ENV), &global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	return v
}
