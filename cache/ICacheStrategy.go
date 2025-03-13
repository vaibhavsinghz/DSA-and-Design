package main

import "container/list"

type ICacheStrategy interface {
	KeyAccessed(c *Cache, val *list.Element)
	Evict(c *Cache)
	Add(c *Cache, obj pair) *list.Element
}
