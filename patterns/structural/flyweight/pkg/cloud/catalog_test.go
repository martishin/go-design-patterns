package cloud_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

func TestNewInstanceTypeCatalog_ReturnsErrDuplicateInstanceType(t *testing.T) {
	first := newInstanceType(t, "standard-small", 2, 8, "x86_64", 12)
	second := newInstanceType(t, "standard-small", 4, 16, "x86_64", 20)

	_, err := cloud.NewInstanceTypeCatalog(first, second)
	if !errors.Is(err, cloud.ErrDuplicateInstanceType) {
		t.Fatalf("got error %v, want %v", err, cloud.ErrDuplicateInstanceType)
	}
}

func TestInstanceTypeCatalog_GetReturnsRegisteredFlyweight(t *testing.T) {
	instanceType := newInstanceType(t, "standard-small", 2, 8, "x86_64", 12)
	catalog, err := cloud.NewInstanceTypeCatalog(instanceType)
	if err != nil {
		t.Fatalf("NewInstanceTypeCatalog() returned error: %v", err)
	}

	first, err := catalog.Get("standard-small")
	if err != nil {
		t.Fatalf("first Get() returned error: %v", err)
	}
	second, err := catalog.Get("standard-small")
	if err != nil {
		t.Fatalf("second Get() returned error: %v", err)
	}

	if first != second {
		t.Fatalf("expected catalog to return the same flyweight")
	}
	if first != instanceType {
		t.Fatalf("expected catalog to return the registered flyweight")
	}
	if catalog.Count() != 1 {
		t.Fatalf("got instance type count %d, want 1", catalog.Count())
	}
}

func TestInstanceTypeCatalog_GetReturnsErrUnknownInstanceType(t *testing.T) {
	catalog, err := cloud.NewInstanceTypeCatalog()
	if err != nil {
		t.Fatalf("NewInstanceTypeCatalog() returned error: %v", err)
	}

	_, err = catalog.Get("missing")
	if !errors.Is(err, cloud.ErrUnknownInstanceType) {
		t.Fatalf("got error %v, want %v", err, cloud.ErrUnknownInstanceType)
	}
}

func newInstanceType(
	t *testing.T,
	name string,
	cpuCount int,
	memoryGB int,
	architecture string,
	hourlyCostCents int,
) *cloud.InstanceType {
	t.Helper()

	instanceType, err := cloud.NewInstanceType(name, cpuCount, memoryGB, architecture, hourlyCostCents)
	if err != nil {
		t.Fatalf("NewInstanceType() returned error: %v", err)
	}

	return instanceType
}
