package cache

import "time"

//Adapter Cache适配器接口
type Adapter interface {
	Set(key string, value []byte, expire time.Duration) error
	//Get Cache Get
	Get(key string) ([]byte, error)
	//Del Cache Del
	Del(key string) (bool, error)
}
