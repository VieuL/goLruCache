Simple lru cache implementation for golang

# Import
```
go get github.com/VieuL/goLruCache
```
```
# Usage
```
lru := NewLRUCache(&LruCacheOptions{
    maxSize: 100000 // max size of cache in bytes
    maxItem: 200 // max number of items in cache
})

lru.Set("key", "value")
lru.Get("key")
lru.Delete("key")
lru.Clear()
```