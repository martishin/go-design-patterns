package cachedregistry

import (
	"errors"
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/pkg/registry"
)

var ErrNilRegistry = errors.New("cachedregistry: nil artifact registry")

type CachedArtifactRegistry struct {
	registry registry.ArtifactRegistry
	cache    map[string]registry.Artifact
}

func NewCachedArtifactRegistry(artifactRegistry registry.ArtifactRegistry) (*CachedArtifactRegistry, error) {
	if artifactRegistry == nil {
		return nil, ErrNilRegistry
	}

	return &CachedArtifactRegistry{
		registry: artifactRegistry,
		cache:    make(map[string]registry.Artifact),
	}, nil
}

func (c *CachedArtifactRegistry) Get(service string, version string) (registry.Artifact, error) {
	key := cacheKey(service, version)
	if artifact, exists := c.cache[key]; exists {
		return artifact, nil
	}

	artifact, err := c.registry.Get(service, version)
	if err != nil {
		return registry.Artifact{}, err
	}

	c.cache[key] = artifact
	return artifact, nil
}

func (c *CachedArtifactRegistry) CachedArtifacts() int {
	return len(c.cache)
}

func cacheKey(service string, version string) string {
	return fmt.Sprintf("%s:%s", service, version)
}
