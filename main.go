package main

import (
	"cache/storage"
)

func main() {

	hms := storage.New[string, int](5)

	hms.Add("5", 5)
	hms.Add("4", 4)

	a, _ := hms.Get("2")
	println(a)
}
