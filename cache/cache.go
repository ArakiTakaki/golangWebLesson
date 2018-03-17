package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheInstance *cache.Cache = nil

func getInstance() *cache.Cache {
	if cacheInstance != nil {
		cacheInstance = cache.New(5*time.Minute, 30*time.Second)
	}
	return cacheInstance
}

func set(key string, value interface{}) {
	cacheInstance.Set(key, value, 0)
}
