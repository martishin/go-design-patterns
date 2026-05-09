package cache

type Entry struct {
	Key        string
	CreatedAt  int
	LastUsedAt int
}

type EvictionStrategy interface {
	Name() string
	Evict(entries []Entry) (string, error)
}
