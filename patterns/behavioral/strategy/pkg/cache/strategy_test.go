package cache_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"
)

type evictionStrategyStub struct {
	evictedKey string
}

func (s evictionStrategyStub) Name() string {
	return "test strategy"
}

func (s evictionStrategyStub) Evict(_ []cache.Entry) (string, error) {
	return s.evictedKey, nil
}

func TestEvictionStrategy_InterfaceCanBeImplemented(t *testing.T) {
	var strategy cache.EvictionStrategy = evictionStrategyStub{evictedKey: "a"}

	if strategy.Name() != "test strategy" {
		t.Fatalf("got strategy name %q, want %q", strategy.Name(), "test strategy")
	}

	evictedKey, err := strategy.Evict([]cache.Entry{{Key: "a"}})
	if err != nil {
		t.Fatalf("Evict() returned error: %v", err)
	}
	if evictedKey != "a" {
		t.Fatalf("got evicted key %q, want %q", evictedKey, "a")
	}
}
