package remoteregistry

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/pkg/registry"
)

type RemoteArtifactRegistry struct {
	host  string
	calls int
}

func NewRemoteArtifactRegistry(host string) *RemoteArtifactRegistry {
	return &RemoteArtifactRegistry{host: host}
}

func (r *RemoteArtifactRegistry) Get(service string, version string) (registry.Artifact, error) {
	if service == "" || version == "" {
		return registry.Artifact{}, registry.ErrInvalidArtifactRequest
	}

	r.calls++

	return registry.Artifact{
		Service: service,
		Version: version,
		Image:   fmt.Sprintf("%s/%s:%s", r.host, service, version),
		Digest:  fmt.Sprintf("sha256:%s-%s", service, version),
		SizeMB:  128,
	}, nil
}

func (r *RemoteArtifactRegistry) Calls() int {
	return r.calls
}
