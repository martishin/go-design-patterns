package cloud_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

func TestNewInstanceType_ReturnsErrInvalidInstanceType(t *testing.T) {
	instanceType, err := cloud.NewInstanceType("standard-small", 0, 8, "x86_64", 12)
	if !errors.Is(err, cloud.ErrInvalidInstanceType) {
		t.Fatalf("got error %v, want %v", err, cloud.ErrInvalidInstanceType)
	}
	if instanceType != nil {
		t.Fatalf("got instance type %v, want nil", instanceType)
	}
}

func TestInstanceType_ReturnsIntrinsicState(t *testing.T) {
	instanceType, err := cloud.NewInstanceType("standard-small", 2, 8, "x86_64", 12)
	if err != nil {
		t.Fatalf("NewInstanceType() returned error: %v", err)
	}

	if instanceType.Name() != "standard-small" {
		t.Fatalf("got name %q, want %q", instanceType.Name(), "standard-small")
	}
	if instanceType.CPUCount() != 2 {
		t.Fatalf("got CPU count %d, want 2", instanceType.CPUCount())
	}
	if instanceType.MemoryGB() != 8 {
		t.Fatalf("got memory %d, want 8", instanceType.MemoryGB())
	}
	if instanceType.Architecture() != "x86_64" {
		t.Fatalf("got architecture %q, want %q", instanceType.Architecture(), "x86_64")
	}
	if instanceType.HourlyCostCents() != 12 {
		t.Fatalf("got hourly cost %d, want 12", instanceType.HourlyCostCents())
	}
}

func TestInstanceType_CostCentsUsesHourlyRate(t *testing.T) {
	instanceType, err := cloud.NewInstanceType("standard-small", 2, 8, "x86_64", 12)
	if err != nil {
		t.Fatalf("NewInstanceType() returned error: %v", err)
	}

	got := instanceType.CostCents(720)
	want := 12 * 720
	if got != want {
		t.Fatalf("got cost %d, want %d", got, want)
	}
}
