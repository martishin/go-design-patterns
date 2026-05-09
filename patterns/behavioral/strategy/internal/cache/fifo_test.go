package cache_test

import (
	"errors"
	"testing"

	internalcache "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/internal/cache"
	cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"
)

func TestFIFO_ImplementsEvictionStrategy(t *testing.T) {
	var strategy cacheapi.EvictionStrategy = internalcache.NewFIFO()

	if strategy.Name() != "FIFO" {
		t.Fatalf("got strategy name %q, want %q", strategy.Name(), "FIFO")
	}
}

func TestFIFO_EvictsOldestCreatedEntry(t *testing.T) {
	strategy := internalcache.NewFIFO()
	entries := []cacheapi.Entry{
		{Key: "b", CreatedAt: 2, LastUsedAt: 10},
		{Key: "a", CreatedAt: 1, LastUsedAt: 20},
	}

	evictedKey, err := strategy.Evict(entries)
	if err != nil {
		t.Fatalf("Evict() returned error: %v", err)
	}

	if evictedKey != "a" {
		t.Fatalf("got evicted key %q, want %q", evictedKey, "a")
	}
}

func TestFIFO_EvictReturnsErrorForEmptyEntries(t *testing.T) {
	strategy := internalcache.NewFIFO()

	_, err := strategy.Evict(nil)

	if !errors.Is(err, internalcache.ErrNoEntries) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrNoEntries)
	}
}

func TestCache_UsesFIFOStrategy(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	if _, err := cache.Put("a", "1"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	if _, err := cache.Put("b", "2"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	evictedKey, err := cache.Put("c", "3")
	if err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	if evictedKey != "a" {
		t.Fatalf("got evicted key %q, want %q", evictedKey, "a")
	}
}
