package main

import (
	"container/list"
	"sync"
)

const cacheSize = 10

type CacheLoader interface {
	Load(key string) string
}

type Node struct {
	Key   string
	Value string
}

type CacheInMemory struct {
	loader CacheLoader
	nodes  *list.List

	lock sync.Mutex
	mp   map[string]*list.Element
}

func NewCacheInMemory(loader CacheLoader) *CacheInMemory {
	return &CacheInMemory{
		loader: loader,
		nodes:  list.New(),
		mp:     make(map[string]*list.Element),
	}
}

func (cache *CacheInMemory) Get(key string) string {
	val, exists := cache.mp[key]
	if exists {
		cache.lock.Lock()
		cache.nodes.MoveToFront(val)
		cache.lock.Unlock()
		return val.Value.(Node).Value
	}
	cache.lock.Lock()
	defer cache.lock.Unlock()
	nodeValue := Node{
		Key:   key,
		Value: cache.loader.Load(key),
	}
	if cache.nodes.Len() == cacheSize {
		lastNode := cache.nodes.Back()
		delete(cache.mp, lastNode.Value.(Node).Key)
		cache.nodes.Remove(lastNode)
	}
	cache.nodes.PushFront(nodeValue)
	cache.mp[key] = cache.nodes.Front()
	return nodeValue.Value
}
