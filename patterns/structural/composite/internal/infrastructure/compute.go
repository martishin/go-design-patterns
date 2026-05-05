package infrastructure

import (
	"crypto/rand"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

const cpuMonthlyCostCents = 10 * 100

type ComputeInstance struct {
	id       string
	name     string
	cpuCount int
}

func NewComputeInstance(name string, cpuCount int) (cost.CostComponent, error) {
	if cpuCount < 0 {
		return nil, cost.ErrInvalidResourceCount
	}

	return &ComputeInstance{
		id:       rand.Text(),
		name:     name,
		cpuCount: cpuCount,
	}, nil
}

func (c *ComputeInstance) MonthlyCostCents() int {
	return c.cpuCount * cpuMonthlyCostCents
}
