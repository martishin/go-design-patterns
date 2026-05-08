package cloud

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicateInstanceType = errors.New("cloud: duplicate instance type")
	ErrUnknownInstanceType   = errors.New("cloud: unknown instance type")
)

type InstanceTypeCatalog struct {
	instanceTypes map[string]*InstanceType
}

func NewInstanceTypeCatalog(instanceTypes ...*InstanceType) (*InstanceTypeCatalog, error) {
	catalog := &InstanceTypeCatalog{
		instanceTypes: make(map[string]*InstanceType, len(instanceTypes)),
	}

	for _, instanceType := range instanceTypes {
		if instanceType == nil {
			return nil, ErrInvalidInstanceType
		}
		if _, exists := catalog.instanceTypes[instanceType.Name()]; exists {
			return nil, fmt.Errorf("%w: %s", ErrDuplicateInstanceType, instanceType.Name())
		}

		catalog.instanceTypes[instanceType.Name()] = instanceType
	}

	return catalog, nil
}

func NewDefaultInstanceTypeCatalog() (*InstanceTypeCatalog, error) {
	standardSmall, err := NewInstanceType("standard-small", 2, 8, "x86_64", 12)
	if err != nil {
		return nil, err
	}
	standardLarge, err := NewInstanceType("standard-large", 8, 32, "x86_64", 42)
	if err != nil {
		return nil, err
	}
	armMedium, err := NewInstanceType("arm-medium", 4, 16, "arm64", 20)
	if err != nil {
		return nil, err
	}

	return NewInstanceTypeCatalog(standardSmall, standardLarge, armMedium)
}

func (c *InstanceTypeCatalog) Get(name string) (*InstanceType, error) {
	instanceType, exists := c.instanceTypes[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrUnknownInstanceType, name)
	}

	return instanceType, nil
}

func (c *InstanceTypeCatalog) Count() int {
	return len(c.instanceTypes)
}
