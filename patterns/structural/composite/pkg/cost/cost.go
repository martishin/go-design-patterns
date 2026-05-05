package cost

import "errors"

var (
	ErrNilComponent         = errors.New("cost: nil component")
	ErrInvalidResourceCount = errors.New("cost: resource count must be non-negative")
)

type CostComponent interface {
	MonthlyCostCents() int
}
