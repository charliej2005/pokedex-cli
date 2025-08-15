package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.RWMutex{},
	}
	return c
}
