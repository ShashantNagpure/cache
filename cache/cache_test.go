package cache

import "testing"

func TestCache(t *testing.T) {

	cache := New[string, int](3)

	cache.put("1", 1)
	cache.put("2", 2)

	cache.put("3", 3)

	println(cache.get("1"))

	cache.put("1", 4)
}
