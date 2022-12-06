package main

import "cache/cache"

func main() {

	cache := cache.New[string, int](3)

	cache.Put("1", 1)
	cache.Put("2", 2)

	v, _ := cache.Get("1")
	println(v)

	cache.Put("3", 3)

	cache.Put("4", 4)

	_, err := cache.Get("2")
	println(err.Error())

}
