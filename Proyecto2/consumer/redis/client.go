package redis

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisLock = &sync.Mutex{}

var redisClient *redis.Client

func Connect2Redis() *redis.Client {

	host := "redis-service"
	port := "6379"
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	return client
}

func GetRedisInstance() *redis.Client {

	if redisClient == nil {
		redisLock.Lock()
		defer redisLock.Unlock()
		if redisClient == nil {
			fmt.Println("Creating single redis instance now.")
			redisClient = Connect2Redis()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return redisClient
}
