package cache_test

import (
	"errors"
	"testing"

	internalcache "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/internal/cache"
	cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"
)

func TestLRU_ImplementsEvictionStrategy(t *testing.T) {
	var strategy cacheapi.EvictionStrategy = internalcache.NewLRU()

	if strategy.Name() != "LRU" {
		t.Fatalf("got strategy name %q, want %q", strategy.Name(), "LRU")
	}
}

func TestLRU_EvictsLeastRecentlyUsedEntry(t *testing.T) {
	strategy := internalcache.NewLRU()
	entries := []cacheapi.Entry{
		{Key: "a", CreatedAt: 1, LastUsedAt: 20},
		{Key: "b", CreatedAt: 2, LastUsedAt: 10},
	}

	evictedKey, err := strategy.Evict(entries)
	if err != nil {
		t.Fatalf("Evict() returned error: %v", err)
	}

	if evictedKey != "b" {
		t.Fatalf("got evicted key %q, want %q", evictedKey, "b")
	}
}

func TestLRU_EvictReturnsErrorForEmptyEntries(t *testing.T) {
	strategy := internalcache.NewLRU()

	_, err := strategy.Evict(nil)

	if !errors.Is(err, internalcache.ErrNoEntries) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrNoEntries)
	}
}

func TestCache_UsesLRUStrategy(t *testing.T) {
	cache := newTestCache(t, internalcache.NewLRU())

	if _, err := cache.Put("a", "1"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	if _, err := cache.Put("b", "2"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	if _, ok := cache.Get("a"); !ok {
		t.Fatal("expected key a to exist")
	}

	evictedKey, err := cache.Put("c", "3")
	if err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	if evictedKey != "b" {
		t.Fatalf("got evicted key %q, want %q", evictedKey, "b")
	}
}
