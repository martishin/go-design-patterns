package cost_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

type componentStub struct {
	monthlyCostCents int
}

func (c componentStub) MonthlyCostCents() int {
	return c.monthlyCostCents
}

func TestNewCostGroup_ReturnsErrNilComponentWhenAnyChildIsNil(t *testing.T) {
	group, err := cost.NewCostGroup("production", componentStub{monthlyCostCents: 1000}, nil)
	if !errors.Is(err, cost.ErrNilComponent) {
		t.Fatalf("got error %v, want %v", err, cost.ErrNilComponent)
	}
	if group != nil {
		t.Fatalf("got group %v, want nil", group)
	}
}

func TestCostGroup_AddReturnsErrNilComponent(t *testing.T) {
	group, err := cost.NewCostGroup("production")
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	err = group.Add(nil)
	if !errors.Is(err, cost.ErrNilComponent) {
		t.Fatalf("got error %v, want %v", err, cost.ErrNilComponent)
	}
}

func TestCostGroup_MonthlyCostCentsReturnsZeroForEmptyGroup(t *testing.T) {
	group, err := cost.NewCostGroup("production")
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	got := group.MonthlyCostCents()
	if got != 0 {
		t.Fatalf("got monthly cost %d, want 0", got)
	}
}

func TestCostGroup_MonthlyCostCentsSumsDirectChildren(t *testing.T) {
	group, err := cost.NewCostGroup(
		"production",
		componentStub{monthlyCostCents: 4000},
		componentStub{monthlyCostCents: 12000},
	)
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	got := group.MonthlyCostCents()
	want := 16000

	if got != want {
		t.Fatalf("got monthly cost %d, want %d", got, want)
	}
}

func TestCostGroup_MonthlyCostCentsSumsNestedGroups(t *testing.T) {
	production, err := cost.NewCostGroup(
		"production",
		componentStub{monthlyCostCents: 4000},
		componentStub{monthlyCostCents: 20000},
	)
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	staging, err := cost.NewCostGroup(
		"staging",
		componentStub{monthlyCostCents: 2000},
		componentStub{monthlyCostCents: 5000},
	)
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	allEnvironments, err := cost.NewCostGroup("all-environments", production, staging)
	if err != nil {
		t.Fatalf("NewCostGroup() returned error: %v", err)
	}

	got := allEnvironments.MonthlyCostCents()
	want := 31000

	if got != want {
		t.Fatalf("got monthly cost %d, want %d", got, want)
	}
}
