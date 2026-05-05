package infrastructure

import (
	"crypto/rand"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

const gbMonthlyCostsCents = 1 * 100

type ManagedDatabase struct {
	id      string
	name    string
	gbCount int
}

func NewManagedDatabase(name string, gbCount int) (cost.CostComponent, error) {
	if gbCount < 0 {
		return nil, cost.ErrInvalidResourceCount
	}

	return &ManagedDatabase{
		id:      rand.Text(),
		name:    name,
		gbCount: gbCount,
	}, nil
}

func (m ManagedDatabase) MonthlyCostCents() int {
	return m.gbCount * gbMonthlyCostsCents
}
