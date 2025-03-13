package main

import "container/list"

type pair struct {
	key, value int
}
type Cache struct {
	Memo          map[int]*list.Element
	List          *list.List
	Size          int
	Capacity      int
	CacheStrategy ICacheStrategy
}

func NewCache(capacity int, evictionAlgo ICacheStrategy) *Cache {
	return &Cache{
		Memo:          map[int]*list.Element{},
		List:          list.New(),
		Capacity:      capacity,
		CacheStrategy: evictionAlgo,
	}
}

func (c *Cache) Get(key int) int {
	val, exist := c.Memo[key]
	if !exist {
		return -1
	}
	c.CacheStrategy.KeyAccessed(c, val)
	return val.Value.(pair).value
}

func (c *Cache) Put(key int, value int) {
	val, exist := c.Memo[key]
	if exist {
		val.Value = pair{key, value}
		c.CacheStrategy.KeyAccessed(c, val)
	} else {
		if c.Size == c.Capacity {
			c.CacheStrategy.Evict(c)
		} else {
			c.Size++
		}
		c.Memo[key] = c.CacheStrategy.Add(c, pair{key, value})
	}
}
