package cost

import (
	"crypto/rand"
)

type CostGroup struct {
	id             string
	name           string
	costComponents []CostComponent
}

func NewCostGroup(name string, components ...CostComponent) (*CostGroup, error) {
	for _, component := range components {
		if component == nil {
			return nil, ErrNilComponent
		}
	}

	return &CostGroup{
		id:             rand.Text(),
		name:           name,
		costComponents: components,
	}, nil
}

func (g *CostGroup) Add(component CostComponent) error {
	if component == nil {
		return ErrNilComponent
	}

	g.costComponents = append(g.costComponents, component)
	return nil
}

func (g *CostGroup) MonthlyCostCents() int {
	resultCost := 0

	for _, component := range g.costComponents {
		resultCost += component.MonthlyCostCents()
	}

	return resultCost
}
