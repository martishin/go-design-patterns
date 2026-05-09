package cache_test

import (
	"errors"
	"reflect"
	"testing"

	internalcache "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/internal/cache"
	cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"
)

type failingStrategy struct{}

func (s failingStrategy) Name() string {
	return "failing"
}

func (s failingStrategy) Evict(_ []cacheapi.Entry) (string, error) {
	return "", internalcache.ErrNoEntries
}

func TestNewCache_ReturnsErrorForInvalidCapacity(t *testing.T) {
	_, err := internalcache.NewCache(0, internalcache.NewFIFO())

	if !errors.Is(err, internalcache.ErrInvalidCapacity) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrInvalidCapacity)
	}
}

func TestNewCache_ReturnsErrorForNilStrategy(t *testing.T) {
	_, err := internalcache.NewCache(2, nil)

	if !errors.Is(err, internalcache.ErrNilStrategy) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrNilStrategy)
	}
}

func TestCache_SetStrategyChangesActiveStrategy(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	if err := cache.SetStrategy(internalcache.NewLRU()); err != nil {
		t.Fatalf("SetStrategy() returned error: %v", err)
	}

	if cache.StrategyName() != "LRU" {
		t.Fatalf("got strategy name %q, want %q", cache.StrategyName(), "LRU")
	}
}

func TestCache_SetStrategyReturnsErrorForNilStrategy(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	err := cache.SetStrategy(nil)

	if !errors.Is(err, internalcache.ErrNilStrategy) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrNilStrategy)
	}
}

func TestCache_PutReturnsErrorForEmptyKey(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	_, err := cache.Put("", "1")

	if !errors.Is(err, internalcache.ErrEmptyKey) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrEmptyKey)
	}
}

func TestCache_PutUpdatesExistingEntry(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	if _, err := cache.Put("a", "1"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	evictedKey, err := cache.Put("a", "2")
	if err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	if evictedKey != "" {
		t.Fatalf("got evicted key %q, want empty", evictedKey)
	}

	value, ok := cache.Get("a")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if value != "2" {
		t.Fatalf("got value %q, want %q", value, "2")
	}
}

func TestCache_PutReturnsStrategyError(t *testing.T) {
	cache := newTestCache(t, failingStrategy{})

	if _, err := cache.Put("a", "1"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	if _, err := cache.Put("b", "2"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	_, err := cache.Put("c", "3")
	if !errors.Is(err, internalcache.ErrNoEntries) {
		t.Fatalf("got error %v, want %v", err, internalcache.ErrNoEntries)
	}
}

func TestCache_GetReturnsFalseForMissingKey(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	_, ok := cache.Get("missing")

	if ok {
		t.Fatal("got ok=true, want false")
	}
}

func TestCache_KeysReturnsSortedKeys(t *testing.T) {
	cache := newTestCache(t, internalcache.NewFIFO())

	if _, err := cache.Put("b", "2"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}
	if _, err := cache.Put("a", "1"); err != nil {
		t.Fatalf("Put() returned error: %v", err)
	}

	want := []string{"a", "b"}
	if !reflect.DeepEqual(cache.Keys(), want) {
		t.Fatalf("got keys %v, want %v", cache.Keys(), want)
	}
}

func newTestCache(t *testing.T, strategy cacheapi.EvictionStrategy) *internalcache.Cache {
	t.Helper()

	cache, err := internalcache.NewCache(2, strategy)
	if err != nil {
		t.Fatalf("NewCache() returned error: %v", err)
	}

	return cache
}
