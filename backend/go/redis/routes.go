package  redis

import (

	"github.com/go-redis/redis"
)

//CacheServerAPI :  Is in charge to provide all the posible transactions to the router
type CacheServerAPI struct {
	redisClient *redis.Client
}

//ProvideCacheAPI :  Provides to the router the entry api
func ProvideCacheAPI(redisClient *redis.Client) CacheServerAPI {
	return CacheServerAPI{redisClient: redisClient}
}


