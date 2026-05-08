package infrastructure

import (
	"errors"
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

var (
	ErrNilInstanceType     = errors.New("infrastructure: nil instance type")
	ErrInvalidRunningHours = errors.New("infrastructure: running hours must be non-negative")
)

type ComputeInstance struct {
	id           string
	region       string
	environment  string
	runningHours int
	instanceType *cloud.InstanceType
}

func NewComputeInstance(
	id string,
	region string,
	environment string,
	runningHours int,
	instanceType *cloud.InstanceType,
) (*ComputeInstance, error) {
	if instanceType == nil {
		return nil, ErrNilInstanceType
	}
	if runningHours < 0 {
		return nil, ErrInvalidRunningHours
	}

	return &ComputeInstance{
		id:           id,
		region:       region,
		environment:  environment,
		runningHours: runningHours,
		instanceType: instanceType,
	}, nil
}

func (c *ComputeInstance) ID() string {
	return c.id
}

func (c *ComputeInstance) Region() string {
	return c.region
}

func (c *ComputeInstance) Environment() string {
	return c.environment
}

func (c *ComputeInstance) InstanceType() *cloud.InstanceType {
	return c.instanceType
}

func (c *ComputeInstance) MonthlyCostCents() int {
	return c.instanceType.CostCents(c.runningHours)
}

func (c *ComputeInstance) Description() string {
	return fmt.Sprintf(
		"%s runs %s in %s/%s",
		c.id,
		c.instanceType.Name(),
		c.environment,
		c.region,
	)
}
