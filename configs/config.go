package configs

import (
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	checkError(err)
	err = viper.Unmarshal(&cfg)
	checkError(err)
	return cfg, err
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
