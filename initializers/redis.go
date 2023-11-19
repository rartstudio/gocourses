package initializers

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)


func ConnectToRedis(config *Config) *redis.Client {
	dsn := fmt.Sprintf("%s:%s", config.REDISHOST, config.REDISPORT)
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		Password: "",
		Username: "",
		DB: 0,
	})

	log.Println("🚀 Connected Successfully to the Redis")

	return client
}
