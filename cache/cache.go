package cache

import (
	"cache/eviction"
	"cache/storage"
	"errors"
	"fmt"
)

type Cache[K any, V any] struct {
	eviction eviction.Eviction[K]
	storage  storage.Storage[K, V]
}

func New[K comparable, V any](capacity int) *Cache[K, V] {

	return &Cache[K, V]{
		eviction: eviction.New[K](),
		storage:  storage.New[K, V](capacity),
	}
}

func (cache *Cache[K, V]) Put(key K, value V) error {

	err := cache.storage.Add(key, value)
	if err != nil {
		evictedKey := cache.eviction.Evict()
		err = cache.storage.Remove(evictedKey)
		if err != nil {
			return errors.New(fmt.Sprintf("Unexpected error %v", err))
		}
		println(fmt.Sprintf("evicted %v", evictedKey))

		return cache.Put(key, value)
	}
	cache.eviction.KeyAccessed(key)
	return nil
}

func (cache *Cache[K, V]) Get(key K) (V, error) {
	value, err := cache.storage.Get(key)
	if err == nil {
		cache.eviction.KeyAccessed(key)
	}
	return value, err
}
