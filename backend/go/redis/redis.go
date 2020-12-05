package redis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

//MainAPi :  Struct of the redis package
type MainAPi struct{}

//NewClient :  Creates a new client to use Redis
func NewClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6380",
		Password: "",
		DB:       0,
	})

	return redisClient
}

//Ping :  Make a ping to redis
func Ping(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

//GetRedis :  Get data stored in redis
func (p *MainAPi) GetRedis(client *redis.Client, key string) (string, error) {
	result, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

//GetDB :  Get data stored in other servers
func (p *MainAPi) GetDB(client *redis.Client, url string, data interface{}) interface{} {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return &data

}