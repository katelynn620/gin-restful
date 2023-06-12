package config

import "github.com/spf13/viper"

type Config struct {
	Port  int    `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	viper.SetDefault("port", "8080")

	err = viper.Unmarshal(&c)

	return
}
