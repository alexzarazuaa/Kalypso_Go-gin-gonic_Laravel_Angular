package common

import (
	"github.com/go-redis/redis/v8"
    "context"

)
var ctx = context.Background()

//S_Users :  Struct of the redis package
type S_Users struct{
	Key   string `json:"key"   binding:"required"`
	Value string `json:"value" binding:"required"`
}

//NewClient :  Creates a new client to use Redis
func NewClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6380",
		Password: "",
		DB:       0,
	})

	return redisClient 
}

func SetUser(key string, value string, client *redis.Client) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}