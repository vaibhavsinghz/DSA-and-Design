package main

import "container/list"

type LRUCache struct{}

func NewLRUCache() *LRUCache {
	return &LRUCache{}
}

func (L *LRUCache) KeyAccessed(c *Cache, val *list.Element) {
	c.List.MoveToFront(val)
}

func (L *LRUCache) Evict(c *Cache) {
	lastElement := c.List.Back()
	c.List.Remove(lastElement)

	cacheKey := lastElement.Value.(pair).key
	delete(c.Memo, cacheKey)
}

func (L *LRUCache) Add(c *Cache, obj pair) *list.Element {
	return c.List.PushFront(obj)
}
