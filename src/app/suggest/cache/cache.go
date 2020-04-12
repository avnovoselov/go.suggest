package cache

import (
	. "example.com/suggest/src/app/suggest/client"
	"log"
)

type Cache struct {
	Size            int
	ClearPercentage int

	storage         map[string]SuggestResult
	storageKeyQueue []string
}

func (cache *Cache) Create() {
	cache.storage = make(map[string]SuggestResult)
	log.Printf("Create suggest cache contains %[1]d elements, %[2]d%% clearing", cache.Size, cache.ClearPercentage)
}

func (cache *Cache) Get(query string) (SuggestResult, bool) {
	result, contains := cache.storage[query]
	return result, contains
}

func (cache *Cache) clear() {
	removeSuggestCount := cache.Size / 100 * cache.ClearPercentage
	for i := 0; i < removeSuggestCount; i++ {
		delete(cache.storage, cache.storageKeyQueue[i])
	}
	cache.storageKeyQueue = cache.storageKeyQueue[removeSuggestCount:]
	log.Printf("Remove %[1]d item from cache", removeSuggestCount)
}

func (cache *Cache) Store(query string, suggestResult SuggestResult) {
	cache.storage[query] = suggestResult
	cache.storageKeyQueue = append(cache.storageKeyQueue, query)
	log.Printf("Append new suggest to cache. Cache storage size: %[1]d", len(cache.storage))

	if len(cache.storageKeyQueue) == cache.Size {
		cache.clear()
	}
}
