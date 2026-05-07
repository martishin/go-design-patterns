package registry

import (
	"errors"
	"fmt"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

var ErrEmptyRegistryHost = errors.New("registry: host is required")

type Registry struct {
	host string
}

func NewRegistry(host string) deploy.ArtifactStore {
	return &Registry{host: host}
}

func (r *Registry) Upload(artifact deploy.Artifact) (deploy.ArtifactRef, error) {
	if strings.TrimSpace(r.host) == "" {
		return deploy.ArtifactRef{}, ErrEmptyRegistryHost
	}

	ref := deploy.ArtifactRef{
		URI: fmt.Sprintf("%s/%s", r.host, artifact.Image),
	}

	fmt.Printf("Uploaded artifact to %s\n", ref.URI)

	return ref, nil
}
