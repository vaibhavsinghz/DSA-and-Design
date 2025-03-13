package main

import "fmt"

func main() {
	cache := NewCache(5, NewLRUCache())
	cache.Put(3, 5)
	cache.Put(4, 7)
	cache.Put(6, 8)
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	cache.Put(2, 8)
	cache.Put(9, 18)
	cache.Put(19, 118)
	fmt.Println(cache.Get(6))
}
