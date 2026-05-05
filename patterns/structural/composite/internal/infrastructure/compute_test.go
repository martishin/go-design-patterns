package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/internal/infrastructure"
	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

func TestNewComputeInstance_ReturnsErrInvalidResourceCount(t *testing.T) {
	instance, err := infrastructure.NewComputeInstance("api-server", -1)
	if !errors.Is(err, cost.ErrInvalidResourceCount) {
		t.Fatalf("got error %v, want %v", err, cost.ErrInvalidResourceCount)
	}
	if instance != nil {
		t.Fatalf("got instance %v, want nil", instance)
	}
}

func TestComputeInstance_MonthlyCostCentsReturnsConfiguredCost(t *testing.T) {
	instance, err := infrastructure.NewComputeInstance("api-server", 4)
	if err != nil {
		t.Fatalf("NewComputeInstance() returned error: %v", err)
	}

	got := instance.MonthlyCostCents()
	want := 4 * 10 * 100

	if got != want {
		t.Fatalf("got monthly cost %d, want %d", got, want)
	}
}
