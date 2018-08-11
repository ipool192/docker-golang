package infrastructures

import (
	config "github.com/spf13/viper"
)

// Configuration - struct for initialize config
type Configuration struct{}

func (c *Configuration) InitConfig() {
	config.SetConfigName("App")
	config.AddConfigPath("./configurations")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
}
