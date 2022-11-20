package goLruCache

import (
	"encoding/binary"

	"github.com/go-playground/validator/v10"
)

type LruCache struct {
	options LruCacheOptions
	cache   map[string]any
	keys    []string
}

type LruCacheOptions struct {
	maxSize int
	maxItem int
}

var validate = validator.New()

func NewLruCache(options LruCacheOptions) *LruCache {
	if validationErr := validate.Struct(&options); validationErr != nil {
		panic(validationErr)
	}

	return &LruCache{
		options: options,
		cache:   make(map[string]any),
		keys:    make([]string, options.maxItem),
	}
}

func (lru *LruCache) Get(key string) any {
	go lru.lastResentUsedUpdate(key)
	return lru.cache[key]
}

func (lru *LruCache) Set(key string, value any) {
	if (lru.size()+binary.Size(value) >= lru.options.maxSize) || len(lru.keys) == lru.options.maxItem {
		lru.removeOldest()
	}
	lru.keys = append(lru.keys, key)
	lru.cache[key] = value
}

func (lru *LruCache) Delete(key string) {
	delete(lru.cache, key)
}

func (lru *LruCache) Clear() {
	lru.cache = make(map[string]any)
}

func (lru *LruCache) size() int {
	return binary.Size(lru.cache)
}

func (lru *LruCache) removeOldest() {
	lastElement := lru.keys[len(lru.keys)-1]
	lru.Delete(lastElement)
	lru.keys = lru.keys[:len(lru.keys)-1]
}

func (lru *LruCache) lastResentUsedUpdate(key string) {
	index := indexOf(key, lru.keys)
	lru.keys = append(lru.keys[:index], lru.keys[index+1:]...)
	lru.keys = append(lru.keys, key)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
