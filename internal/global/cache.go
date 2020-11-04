package global

import "github.com/eden-framework/plugin-cache/cache"

var CacheConfig = struct {
	Cache *cache.Cache
}{
	Cache: &cache.Cache{
		Driver: cache.CACHE_DRIVER__REDIS,
		Prefix: "IDP-",
		Host:   "localhost",
		Port:   6379,
		DB:     0,
	},
}
