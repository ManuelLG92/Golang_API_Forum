package storage

import (
	"fmt"
	postDomain "forum/posts/domain"
	userDomain "forum/user/domain"
)

type cacheable interface {
	userDomain.User | postDomain.Post
}

type cache[T cacheable] struct {
	Items map[string]T
}

func NewCache[T cacheable]() *cache[T] {
	c := cache[T]{}
	c.Items = make(map[string]T)
	return &c
}

func (cache *cache[T]) Set(key string, item T) {
	fmt.Println(cache.Items)
	cache.Items[key] = item
	cache.logBuilder("Set", key)
	fmt.Println("cache Items", cache.Items)
}

func (cache *cache[T]) Get(key string) (v *T) {
	if v, ok := cache.Items[key]; ok {
		cache.logBuilder("Found", key)
		return &v
	}
	cache.logBuilder("Not found", key)

	return nil
}

func (cache *cache[T]) Delete(key string) {
	if value := cache.Get(key); value != nil {
		cache.logBuilder("Deleted", key)
		delete(cache.Items, key)
	}
}

func (cache *cache[T]) Reset() {
	cache.Items = make(map[string]T)
}

func (cache *cache[T]) logBuilder(prefix, key string) {
	fmt.Printf("%v item in cache by key: %v and type: %v", prefix, key, *new(T))

}
