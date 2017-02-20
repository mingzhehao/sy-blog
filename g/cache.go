package g

import (
	"time"
)

const (
	blogPrefix    = "b_"
	catalogPrefix = "c_"
	chatPrefix    = "chat_"
	viewPrefix    = "view_"
	CacheExpire   = 3600
)

func BlogCachePut(key string, val interface{}) error {
	return Cache.Put(blogPrefix+key, val, CacheExpire*time.Second)
}

func CatalogCachePut(key string, val interface{}) error {
	return Cache.Put(catalogPrefix+key, val, CacheExpire*time.Second)
}

func BlogCacheGet(key string) interface{} {
	return Cache.Get(blogPrefix + key)
}

func CatalogCacheGet(key string) interface{} {
	return Cache.Get(catalogPrefix + key)
}

func CatalogCacheDel(key string) error {
	return Cache.Delete(catalogPrefix + key)
}

func BlogCacheDel(key string) error {
	return Cache.Delete(blogPrefix + key)
}

/*chat cache*/
func ChatCachePut(key string, val interface{}) error {
	return Cache.Put(chatPrefix+key, val, CacheExpire*time.Second)
}
func ChatCacheGet(key string) interface{} {
	return Cache.Get(chatPrefix + key)
}

/*chat cache*/

/*view cache*/
func ViewCachePut(key string, val interface{}) error {
	return Cache.Put(viewPrefix+key, val, CacheExpire*time.Second)
}
func ViewCacheGet(key string) interface{} {
	return Cache.Get(viewPrefix + key)
}

/*chat cache*/
