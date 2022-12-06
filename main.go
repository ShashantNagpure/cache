package main

import "cache/cache"

func main() {

	cache := cache.New[string, int](3)

	cache.Put("1", 1)
	cache.Put("2", 2)
	cache.Put("3", 3)

	v, _ := cache.Get("1")
	println(v)

	cache.Put("4", 4)

	_, err := cache.Get("2")
	println(err.Error())

	cache.Put("1", 11)
	v, _ = cache.Get("1")
	println(v)

	cache.Put("5", 1)

}
