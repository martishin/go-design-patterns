package cache

import cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"

type FIFO struct{}

func NewFIFO() *FIFO {
	return &FIFO{}
}

func (s *FIFO) Name() string {
	return "FIFO"
}

func (s *FIFO) Evict(entries []cacheapi.Entry) (string, error) {
	if len(entries) == 0 {
		return "", ErrNoEntries
	}

	evictedEntry := entries[0]
	for _, entry := range entries[1:] {
		if entry.CreatedAt < evictedEntry.CreatedAt ||
			(entry.CreatedAt == evictedEntry.CreatedAt && entry.Key < evictedEntry.Key) {
			evictedEntry = entry
		}
	}

	return evictedEntry.Key, nil
}
