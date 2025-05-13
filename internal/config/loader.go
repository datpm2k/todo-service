package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := &AppConfigStruct{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config: %s \n", err))
	}

	AppConfig = config
}
