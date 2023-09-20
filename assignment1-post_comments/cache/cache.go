package cache

import (
	"posts/types"
	"sync"
)

var cache = make(map[string]*[]types.Post)
var mutex = &sync.RWMutex{}

//type CacheStruct struct {
//	mu    *sync.RWMutex
//	cache map[string]string
//}

//func CacheFunc() *CacheStruct {
//	return &CacheStruct{
//		mu:    &sync.RWMutex{},
//		cache: make(map[string]string),
//	}
//}

// set value to cache
// func (c *CacheStruct) Set(key, value string) {
func Set(key string, value *[]types.Post) {
	mutex.Lock()
	defer mutex.Unlock()
	cache[key] = value
}

// Get a value from the cache
// func (c *CacheStruct) Get(key string) (string, bool) {
func Get(key string) (*[]types.Post, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	val, ok := cache[key]
	return val, ok
}

// Delete a key from the cache
//func (c *CacheStruct) Delete(key string) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	delete(c.cache, key)
//}
