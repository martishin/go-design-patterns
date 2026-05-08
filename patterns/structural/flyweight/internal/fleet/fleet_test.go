package fleet_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/internal/fleet"
	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

func TestNewFleet_ReturnsErrNilCatalog(t *testing.T) {
	productionFleet, err := fleet.NewFleet(nil)
	if !errors.Is(err, fleet.ErrNilCatalog) {
		t.Fatalf("got error %v, want %v", err, fleet.ErrNilCatalog)
	}
	if productionFleet != nil {
		t.Fatalf("got fleet %v, want nil", productionFleet)
	}
}

func TestFleet_LaunchReusesInstanceTypeFlyweights(t *testing.T) {
	productionFleet := newFleet(t)

	first, err := productionFleet.Launch("api-1", "eu-west-1", "production", "standard-large", 720)
	if err != nil {
		t.Fatalf("first Launch() returned error: %v", err)
	}
	second, err := productionFleet.Launch("api-2", "eu-west-1", "production", "standard-large", 720)
	if err != nil {
		t.Fatalf("second Launch() returned error: %v", err)
	}

	if first.InstanceType() != second.InstanceType() {
		t.Fatalf("expected instances to share the same instance type flyweight")
	}
	if productionFleet.SharedInstanceTypeCount() != 1 {
		t.Fatalf("got shared instance types %d, want 1", productionFleet.SharedInstanceTypeCount())
	}
	if len(productionFleet.Instances()) != 2 {
		t.Fatalf("got instances %d, want 2", len(productionFleet.Instances()))
	}
}

func TestFleet_TotalMonthlyCostCentsSumsInstances(t *testing.T) {
	productionFleet := newFleet(t)

	if _, err := productionFleet.Launch("api-1", "eu-west-1", "production", "standard-large", 720); err != nil {
		t.Fatalf("Launch() returned error: %v", err)
	}
	if _, err := productionFleet.Launch("worker-1", "eu-west-1", "production", "arm-medium", 360); err != nil {
		t.Fatalf("Launch() returned error: %v", err)
	}

	if productionFleet.SharedInstanceTypeCount() != 2 {
		t.Fatalf("got shared instance types %d, want 2", productionFleet.SharedInstanceTypeCount())
	}

	got := productionFleet.TotalMonthlyCostCents()
	want := 42*720 + 20*360
	if got != want {
		t.Fatalf("got total cost %d, want %d", got, want)
	}
}

func TestFleet_LaunchReturnsErrUnknownInstanceType(t *testing.T) {
	productionFleet := newFleet(t)

	_, err := productionFleet.Launch("api-1", "eu-west-1", "production", "missing", 720)
	if !errors.Is(err, cloud.ErrUnknownInstanceType) {
		t.Fatalf("got error %v, want %v", err, cloud.ErrUnknownInstanceType)
	}
}

func newFleet(t *testing.T) *fleet.Fleet {
	t.Helper()

	catalog, err := cloud.NewDefaultInstanceTypeCatalog()
	if err != nil {
		t.Fatalf("NewDefaultInstanceTypeCatalog() returned error: %v", err)
	}

	productionFleet, err := fleet.NewFleet(catalog)
	if err != nil {
		t.Fatalf("NewFleet() returned error: %v", err)
	}

	return productionFleet
}
