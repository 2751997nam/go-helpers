package libs

import (
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisInstance *redis.Client
var lock = &sync.Mutex{}

func GetRedis() *redis.Client {
	if redisInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if redisInstance == nil {
			rdb := redis.NewClient(&redis.Options{
				Addr:     os.Getenv("REDIS_URI"),
				Password: "", // no password set
				DB:       0,  // use default DB
			})
			redisInstance = rdb
		}
	}

	return redisInstance
}
