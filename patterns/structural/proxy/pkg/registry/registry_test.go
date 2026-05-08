package registry_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/pkg/registry"
)

type registryStub struct{}

func (r registryStub) Get(service string, version string) (registry.Artifact, error) {
	return registry.Artifact{
		Service: service,
		Version: version,
		Image:   service + ":" + version,
	}, nil
}

func TestArtifactRegistry_InterfaceCanBeImplemented(t *testing.T) {
	var artifactRegistry registry.ArtifactRegistry = registryStub{}

	artifact, err := artifactRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		t.Fatalf("Get() returned error: %v", err)
	}
	if artifact.Image != "payments-api:v1.4.2" {
		t.Fatalf("got image %q, want %q", artifact.Image, "payments-api:v1.4.2")
	}
}
