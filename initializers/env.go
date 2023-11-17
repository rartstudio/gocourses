package initializers

import "github.com/spf13/viper"

type Config struct {
	APPNAME string `mapstructure:"appname"`
	PORT int `mapstructure:"port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)

	return
}