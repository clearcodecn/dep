package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func SetConfig(v interface{}) error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension i
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(v)
	})
	go viper.WatchConfig()

	return viper.Unmarshal(v)
}
