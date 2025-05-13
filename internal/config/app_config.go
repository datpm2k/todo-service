package config

var AppConfig *AppConfigStruct

type AppConfigStruct struct {
	Port int    `mapstructure:"port"`
	ENV  string `mapstructure:"env"`
}
