package cachedregistry_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/internal/cachedregistry"
	"github.com/martishin/go-design-patterns/patterns/structural/proxy/pkg/registry"
)

type registrySpy struct {
	artifact registry.Artifact
	err      error
	calls    int
}

func (s *registrySpy) Get(service string, version string) (registry.Artifact, error) {
	s.calls++
	if s.err != nil {
		return registry.Artifact{}, s.err
	}

	return s.artifact, nil
}

func TestNewCachedArtifactRegistry_ReturnsErrNilRegistry(t *testing.T) {
	cachedRegistry, err := cachedregistry.NewCachedArtifactRegistry(nil)
	if !errors.Is(err, cachedregistry.ErrNilRegistry) {
		t.Fatalf("got error %v, want %v", err, cachedregistry.ErrNilRegistry)
	}
	if cachedRegistry != nil {
		t.Fatalf("got registry %v, want nil", cachedRegistry)
	}
}

func TestCachedArtifactRegistry_GetDelegatesOnCacheMiss(t *testing.T) {
	want := registry.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "registry.acme.io/payments-api:v1.4.2",
		Digest:  "sha256:payments-api-v1.4.2",
		SizeMB:  128,
	}
	spy := &registrySpy{artifact: want}
	cachedRegistry, err := cachedregistry.NewCachedArtifactRegistry(spy)
	if err != nil {
		t.Fatalf("NewCachedArtifactRegistry() returned error: %v", err)
	}

	got, err := cachedRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		t.Fatalf("Get() returned error: %v", err)
	}
	if got != want {
		t.Fatalf("got artifact %+v, want %+v", got, want)
	}
	if spy.calls != 1 {
		t.Fatalf("got calls %d, want 1", spy.calls)
	}
	if cachedRegistry.CachedArtifacts() != 1 {
		t.Fatalf("got cached artifacts %d, want 1", cachedRegistry.CachedArtifacts())
	}
}

func TestCachedArtifactRegistry_GetReturnsCachedArtifact(t *testing.T) {
	want := registry.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "registry.acme.io/payments-api:v1.4.2",
		Digest:  "sha256:payments-api-v1.4.2",
		SizeMB:  128,
	}
	spy := &registrySpy{artifact: want}
	cachedRegistry, err := cachedregistry.NewCachedArtifactRegistry(spy)
	if err != nil {
		t.Fatalf("NewCachedArtifactRegistry() returned error: %v", err)
	}

	first, err := cachedRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		t.Fatalf("first Get() returned error: %v", err)
	}
	second, err := cachedRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		t.Fatalf("second Get() returned error: %v", err)
	}

	if first != second {
		t.Fatalf("got second artifact %+v, want %+v", second, first)
	}
	if spy.calls != 1 {
		t.Fatalf("got calls %d, want 1", spy.calls)
	}
}

func TestCachedArtifactRegistry_GetDoesNotCacheErrors(t *testing.T) {
	wantErr := errors.New("remote failed")
	spy := &registrySpy{err: wantErr}
	cachedRegistry, err := cachedregistry.NewCachedArtifactRegistry(spy)
	if err != nil {
		t.Fatalf("NewCachedArtifactRegistry() returned error: %v", err)
	}

	_, err = cachedRegistry.Get("payments-api", "v1.4.2")
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
	if cachedRegistry.CachedArtifacts() != 0 {
		t.Fatalf("got cached artifacts %d, want 0", cachedRegistry.CachedArtifacts())
	}
}
