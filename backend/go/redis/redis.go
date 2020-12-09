package redis

import (

	// "github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

//S_Users :  Struct of the redis package
type S_Users struct{}

//NewClient :  Creates a new client to use Redis
func NewClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6380",
		Password: "",
		DB:       0,
	})

	return redisClient 
}



//GetRedis :  Get data stored in redis
func (p *S_Users) GetRedis(client *redis.Client, key string) (string, error) {
	result, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
