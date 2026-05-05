package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/internal/infrastructure"
	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

func TestNewManagedDatabase_ReturnsErrInvalidResourceCount(t *testing.T) {
	database, err := infrastructure.NewManagedDatabase("postgres", -1)
	if !errors.Is(err, cost.ErrInvalidResourceCount) {
		t.Fatalf("got error %v, want %v", err, cost.ErrInvalidResourceCount)
	}
	if database != nil {
		t.Fatalf("got database %v, want nil", database)
	}
}

func TestManagedDatabase_MonthlyCostCentsReturnsConfiguredCost(t *testing.T) {
	database, err := infrastructure.NewManagedDatabase("postgres", 200)
	if err != nil {
		t.Fatalf("NewManagedDatabase() returned error: %v", err)
	}

	got := database.MonthlyCostCents()
	want := 200 * 100

	if got != want {
		t.Fatalf("got monthly cost %d, want %d", got, want)
	}
}
