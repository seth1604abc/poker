package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Database databaseConfig
}

type databaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"databasename"`
}

var AppConfig config

func InitConfig(env string) {
	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to unmarshal, %v", err)
	}
}