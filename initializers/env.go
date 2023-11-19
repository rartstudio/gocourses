package initializers

import "github.com/spf13/viper"

type Config struct {
	APPNAME string `mapstructure:"appname"`
	PORT int `mapstructure:"port"`
	DBHOST string `mapstructure:"dbhost"`
	DBPORT string `mapstructure:"dbport"`
	DBUSER string `mapstructure:"dbuser"`
	DBPASSWORD string `mapstructure:"dbpassword"`
	DBNAME string `mapstructure:"dbname"`
	JWTSECRET string `mapstructure:"jwtsecret"`
	JWTEXPIRED int `mapstructure:"jwtexpired"`
	EMAILHOST   string `mapstructure:"emailhost"`
	EMAILPORT   int `mapstructure:"emailport"`
	EMAILUSERNAME string `mapstructure:"emailusername"`
	EMAILPASSWORD string `mapstructure:"emailpassword"`
	REDISHOST string `mapstructure:"redishost"`
	REDISPORT string `mapstructure:"redisport"`
	REDISUSERNAME string `mapstructure:"redisusername"`
	REDISPASSWORD string `mapstructure:"redispassword"`
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