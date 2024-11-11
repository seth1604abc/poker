package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Database	databaseConfig	`mapstructure:"database"`
	JWT			jwtConfig		`mapstructure:"jwt"`
}

type databaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"databasename"`
}

type jwtConfig struct {
	SecretKey	string `mapstructure:"secretKey"`
	ExpiredTime	int `mapstructure:"expiredTime"`
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