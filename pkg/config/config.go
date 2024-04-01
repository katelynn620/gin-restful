package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool   `mapstructure:"DEBUG"`
	Port  int    `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	viper.SetDefault("port", "8080")
	viper.SetDefault("debug", "false")

	err = viper.Unmarshal(&c)

	return
}

func GetConfig() (c Config) {
	// load config if not loaded
	if !viper.IsSet("debug") {
		LoadConfig()
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	return
}
