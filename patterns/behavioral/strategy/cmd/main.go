package main

import (
	"fmt"
	"log"

	internalcache "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/internal/cache"
)

func main() {
	cache, err := internalcache.NewCache(2, internalcache.NewFIFO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initial strategy: %s\n", cache.StrategyName())
	put(cache, "a", "1")
	put(cache, "b", "2")
	put(cache, "c", "3")
	fmt.Printf("Cache keys: %v\n", cache.Keys())

	if err := cache.SetStrategy(internalcache.NewLRU()); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Printf("Switch strategy: %s\n", cache.StrategyName())
	get(cache, "b")
	put(cache, "d", "4")
	fmt.Printf("Cache keys: %v\n", cache.Keys())
}

func put(cache *internalcache.Cache, key string, value string) {
	fmt.Printf("Put %s=%s\n", key, value)

	evictedKey, err := cache.Put(key, value)
	if err != nil {
		log.Fatal(err)
	}
	if evictedKey != "" {
		fmt.Printf("Evicted by %s: %s\n", cache.StrategyName(), evictedKey)
	}
}

func get(cache *internalcache.Cache, key string) {
	value, ok := cache.Get(key)
	if !ok {
		fmt.Printf("Get %s: miss\n", key)
		return
	}

	fmt.Printf("Get %s: %s\n", key, value)
}
