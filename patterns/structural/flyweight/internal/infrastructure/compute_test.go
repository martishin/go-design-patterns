package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/internal/infrastructure"
	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

func TestNewComputeInstance_ReturnsErrNilInstanceType(t *testing.T) {
	instance, err := infrastructure.NewComputeInstance("api-1", "eu-west-1", "production", 720, nil)
	if !errors.Is(err, infrastructure.ErrNilInstanceType) {
		t.Fatalf("got error %v, want %v", err, infrastructure.ErrNilInstanceType)
	}
	if instance != nil {
		t.Fatalf("got instance %v, want nil", instance)
	}
}

func TestNewComputeInstance_ReturnsErrInvalidRunningHours(t *testing.T) {
	instanceType := newInstanceType(t)

	instance, err := infrastructure.NewComputeInstance("api-1", "eu-west-1", "production", -1, instanceType)
	if !errors.Is(err, infrastructure.ErrInvalidRunningHours) {
		t.Fatalf("got error %v, want %v", err, infrastructure.ErrInvalidRunningHours)
	}
	if instance != nil {
		t.Fatalf("got instance %v, want nil", instance)
	}
}

func TestComputeInstance_MonthlyCostCentsUsesSharedInstanceType(t *testing.T) {
	instanceType := newInstanceType(t)
	instance, err := infrastructure.NewComputeInstance("api-1", "eu-west-1", "production", 720, instanceType)
	if err != nil {
		t.Fatalf("NewComputeInstance() returned error: %v", err)
	}

	got := instance.MonthlyCostCents()
	want := 12 * 720
	if got != want {
		t.Fatalf("got monthly cost %d, want %d", got, want)
	}
	if instance.InstanceType() != instanceType {
		t.Fatalf("expected compute instance to reference the shared instance type")
	}
}

func TestComputeInstance_DescriptionIncludesContextAndInstanceType(t *testing.T) {
	instanceType := newInstanceType(t)
	instance, err := infrastructure.NewComputeInstance("api-1", "eu-west-1", "production", 720, instanceType)
	if err != nil {
		t.Fatalf("NewComputeInstance() returned error: %v", err)
	}

	got := instance.Description()
	want := "api-1 runs standard-small in production/eu-west-1"
	if got != want {
		t.Fatalf("got description %q, want %q", got, want)
	}
}

func newInstanceType(t *testing.T) *cloud.InstanceType {
	t.Helper()

	instanceType, err := cloud.NewInstanceType("standard-small", 2, 8, "x86_64", 12)
	if err != nil {
		t.Fatalf("NewInstanceType() returned error: %v", err)
	}

	return instanceType
}
