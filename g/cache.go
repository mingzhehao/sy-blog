package g

import (
	"time"
)

const (
	blogPrefix      = "b_"
	catalogPrefix   = "c_"
	chatPrefix      = "chat_"
	chatCacheExpire = 3600
)

func BlogCachePut(key string, val interface{}) error {
	return Cache.Put(blogPrefix+key, val, time.Duration(blogCacheExpire))
}

func CatalogCachePut(key string, val interface{}) error {
	return Cache.Put(catalogPrefix+key, val, time.Duration(catalogCacheExpire))
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
	return Cache.Put(chatPrefix+key, val, time.Duration(chatCacheExpire))
}
func ChatCacheGet(key string) interface{} {
	return Cache.Get(chatPrefix + key)
}

/*chat cache*/
