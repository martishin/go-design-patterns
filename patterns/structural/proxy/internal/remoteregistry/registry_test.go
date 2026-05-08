package remoteregistry_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/internal/remoteregistry"
	"github.com/martishin/go-design-patterns/patterns/structural/proxy/pkg/registry"
)

func TestRemoteArtifactRegistry_GetReturnsErrInvalidArtifactRequest(t *testing.T) {
	remoteRegistry := remoteregistry.NewRemoteArtifactRegistry("registry.acme.io")

	_, err := remoteRegistry.Get("", "v1.4.2")
	if !errors.Is(err, registry.ErrInvalidArtifactRequest) {
		t.Fatalf("got error %v, want %v", err, registry.ErrInvalidArtifactRequest)
	}
}

func TestRemoteArtifactRegistry_GetReturnsArtifact(t *testing.T) {
	remoteRegistry := remoteregistry.NewRemoteArtifactRegistry("registry.acme.io")

	got, err := remoteRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		t.Fatalf("Get() returned error: %v", err)
	}

	want := registry.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "registry.acme.io/payments-api:v1.4.2",
		Digest:  "sha256:payments-api-v1.4.2",
		SizeMB:  128,
	}
	if got != want {
		t.Fatalf("got artifact %+v, want %+v", got, want)
	}
	if remoteRegistry.Calls() != 1 {
		t.Fatalf("got calls %d, want 1", remoteRegistry.Calls())
	}
}
