package storage

import (
	"errors"
	"fmt"
)

type MapBasedStorage[K comparable, V any] struct {
	mapper   map[K]V
	capacity int
}

func New[K comparable, V any](capacity int) *MapBasedStorage[K, V] {

	return &MapBasedStorage[K, V]{
		mapper:   make(map[K]V),
		capacity: capacity,
	}
}

func (ms *MapBasedStorage[K, V]) Add(key K, value V) error {

	if ms.isStorageFull() {
		return errors.New("capacity full")
	}
	ms.mapper[key] = value
	return nil
}

func (ms *MapBasedStorage[K, V]) isStorageFull() bool {

	return ms.capacity == len(ms.mapper)

}

func (ms *MapBasedStorage[K, V]) Remove(key K) error {

	if _, ok := ms.mapper[key]; ok {
		delete(ms.mapper, key)
		return nil
	} else {
		return errors.New(fmt.Sprintf("key not found: %v", key))
	}

}

func (ms *MapBasedStorage[K, V]) Get(key K) (V, error) {

	var v V
	if val, ok := ms.mapper[key]; ok {
		v = val
		return v, nil
	} else {
		return v, errors.New(fmt.Sprintf("key not found: %v", key))
	}

}
