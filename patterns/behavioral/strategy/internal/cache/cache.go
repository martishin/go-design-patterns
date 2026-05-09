package cache

import (
	"errors"
	"sort"

	cacheapi "github.com/martishin/go-design-patterns/patterns/behavioral/strategy/pkg/cache"
)

var (
	ErrInvalidCapacity = errors.New("cache: capacity must be positive")
	ErrNilStrategy     = errors.New("cache: nil eviction strategy")
	ErrEmptyKey        = errors.New("cache: key is empty")
	ErrNoEntries       = errors.New("cache: no entries to evict")
)

type Cache struct {
	capacity int
	strategy cacheapi.EvictionStrategy
	entries  map[string]entry
	clock    int
}

type entry struct {
	value      string
	createdAt  int
	lastUsedAt int
}

func NewCache(capacity int, strategy cacheapi.EvictionStrategy) (*Cache, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}
	if strategy == nil {
		return nil, ErrNilStrategy
	}

	return &Cache{
		capacity: capacity,
		strategy: strategy,
		entries:  make(map[string]entry),
	}, nil
}

func (c *Cache) SetStrategy(strategy cacheapi.EvictionStrategy) error {
	if strategy == nil {
		return ErrNilStrategy
	}

	c.strategy = strategy
	return nil
}

func (c *Cache) StrategyName() string {
	return c.strategy.Name()
}

func (c *Cache) Put(key string, value string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	c.tick()

	if existingEntry, ok := c.entries[key]; ok {
		existingEntry.value = value
		existingEntry.lastUsedAt = c.clock
		c.entries[key] = existingEntry
		return "", nil
	}

	evictedKey := ""
	if len(c.entries) >= c.capacity {
		keyToEvict, err := c.strategy.Evict(c.snapshot())
		if err != nil {
			return "", err
		}

		delete(c.entries, keyToEvict)
		evictedKey = keyToEvict
	}

	c.entries[key] = entry{
		value:      value,
		createdAt:  c.clock,
		lastUsedAt: c.clock,
	}

	return evictedKey, nil
}

func (c *Cache) Get(key string) (string, bool) {
	existingEntry, ok := c.entries[key]
	if !ok {
		return "", false
	}

	c.tick()
	existingEntry.lastUsedAt = c.clock
	c.entries[key] = existingEntry

	return existingEntry.value, true
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.entries))
	for key := range c.entries {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (c *Cache) snapshot() []cacheapi.Entry {
	entries := make([]cacheapi.Entry, 0, len(c.entries))
	for key, entry := range c.entries {
		entries = append(entries, cacheapi.Entry{
			Key:        key,
			CreatedAt:  entry.createdAt,
			LastUsedAt: entry.lastUsedAt,
		})
	}

	return entries
}

func (c *Cache) tick() {
	c.clock++
}
