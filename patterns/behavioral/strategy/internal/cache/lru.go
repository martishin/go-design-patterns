package cache

import cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"

type LRU struct{}

func NewLRU() *LRU {
	return &LRU{}
}

func (s *LRU) Name() string {
	return "LRU"
}

func (s *LRU) Evict(entries []cacheapi.Entry) (string, error) {
	if len(entries) == 0 {
		return "", ErrNoEntries
	}

	evictedEntry := entries[0]
	for _, entry := range entries[1:] {
		if entry.LastUsedAt < evictedEntry.LastUsedAt ||
			(entry.LastUsedAt == evictedEntry.LastUsedAt && entry.Key < evictedEntry.Key) {
			evictedEntry = entry
		}
	}

	return evictedEntry.Key, nil
}
