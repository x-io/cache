package cache

import "time"

var (
	cache Adapter
)

//Init Init
func Init(adapter Adapter) {
	cache = adapter
}

//Set Cache Set
func Set(key string, value []byte, expire time.Duration) error {
	return cache.Set(key, value, expire)
}

//Get Cache Get
func Get(key string) ([]byte, error) {
	return cache.Get(key)
}

//Del Cache Del
func Del(key string) (bool, error) {
	return cache.Del(key)
}
