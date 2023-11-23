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
	S3ACCESSKEY string `mapstructure:"s3Accesskey"`
	S3SECRETKEY string `mapstructure:"s3Secretkey"`
	S3REGION    string `mapstructure:"s3Region"`
	S3BUCKET    string `mapstructure:"s3Bucket"`
	S3ENDPOINT  string `mapstructure:"s3Endpoint"`
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