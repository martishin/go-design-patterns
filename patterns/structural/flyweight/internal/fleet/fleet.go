package fleet

import (
	"errors"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/internal/infrastructure"
	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

var ErrNilCatalog = errors.New("fleet: nil instance type catalog")

type Fleet struct {
	catalog   *cloud.InstanceTypeCatalog
	instances []*infrastructure.ComputeInstance
}

func NewFleet(catalog *cloud.InstanceTypeCatalog) (*Fleet, error) {
	if catalog == nil {
		return nil, ErrNilCatalog
	}

	return &Fleet{
		catalog: catalog,
	}, nil
}

func (f *Fleet) Launch(
	id string,
	region string,
	environment string,
	typeName string,
	runningHours int,
) (*infrastructure.ComputeInstance, error) {
	instanceType, err := f.catalog.Get(typeName)
	if err != nil {
		return nil, err
	}

	instance, err := infrastructure.NewComputeInstance(
		id,
		region,
		environment,
		runningHours,
		instanceType,
	)
	if err != nil {
		return nil, err
	}

	f.instances = append(f.instances, instance)
	return instance, nil
}

func (f *Fleet) Instances() []*infrastructure.ComputeInstance {
	return append([]*infrastructure.ComputeInstance(nil), f.instances...)
}

func (f *Fleet) TotalMonthlyCostCents() int {
	total := 0
	for _, instance := range f.instances {
		total += instance.MonthlyCostCents()
	}

	return total
}

func (f *Fleet) SharedInstanceTypeCount() int {
	seen := make(map[*cloud.InstanceType]struct{})
	for _, instance := range f.instances {
		seen[instance.InstanceType()] = struct{}{}
	}

	return len(seen)
}
