package cache

import (
	"errors"
	"time"
)

var (
	cache Adapter
)

var NotFound = errors.New("Not Found")

//Init Init
func Init(adapter Adapter) error {
	cache = adapter
	return nil
}

//Set Cache Set
func Set(key string, value []byte, expiration time.Duration) error {
	return cache.Set(key, value, expiration)
}

//Get Cache Get
func Get(key string) ([]byte, error) {
	return cache.Get(key)
}

//Del Cache Del
func Del(key string) error {
	return cache.Del(key)
}
