package registry_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/registry"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func TestRegistry_UploadReturnsArtifactRef(t *testing.T) {
	artifactStore := registry.NewRegistry("registry.acme.io")

	got, err := artifactStore.Upload(deploy.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "payments-api:v1.4.2",
	})
	if err != nil {
		t.Fatalf("Upload() returned error: %v", err)
	}

	want := deploy.ArtifactRef{URI: "registry.acme.io/payments-api:v1.4.2"}
	if got != want {
		t.Fatalf("got artifact ref %+v, want %+v", got, want)
	}
}

func TestRegistry_UploadReturnsErrEmptyRegistryHost(t *testing.T) {
	artifactStore := registry.NewRegistry("")

	_, err := artifactStore.Upload(deploy.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "payments-api:v1.4.2",
	})
	if !errors.Is(err, registry.ErrEmptyRegistryHost) {
		t.Fatalf("got error %v, want %v", err, registry.ErrEmptyRegistryHost)
	}
}
