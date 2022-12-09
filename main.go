package main

import (
	"cache/cache"
	"cache/eviction"
)

func main() {
	TestLRU()
	TestLFU()

}

func TestLRU() {
	cache := cache.New[string, int](3, eviction.NewLRUPolicy[string]())

	cache.Put("1", 1)
	cache.Put("2", 2)
	cache.Put("3", 3)

	v, _ := cache.Get("1")
	println(v)

	cache.Put("4", 4)

	_, err := cache.Get("2")
	println(err.Error())

	cache.Put("1", 11)

	cache.Put("5", 1)

	v, _ = cache.Get("1")
	println(v)

	cache.Put("2", 24)

	_, err = cache.Get("4")
	println(err.Error())

	v, _ = cache.Get("2")
	println(v)

	cache.Put("4", 44)
}

func TestLFU() {

	cache := cache.New[string, int](3, eviction.NewLFUPolicy[string]())

	cache.Put("1", 1)
	cache.Put("2", 2)
	cache.Get("1")
	cache.Put("3", 1)
	cache.Get("2")
	cache.Put("4", 4)
	cache.Get("4")
	cache.Get("4")
	cache.Get("1")
	cache.Put("5", 3)

}
