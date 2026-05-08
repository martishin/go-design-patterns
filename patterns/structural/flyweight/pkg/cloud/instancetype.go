package cloud

import "errors"

var ErrInvalidInstanceType = errors.New("cloud: invalid instance type")

type InstanceType struct {
	name            string
	cpuCount        int
	memoryGB        int
	architecture    string
	hourlyCostCents int
}

func NewInstanceType(name string, cpuCount int, memoryGB int, architecture string, hourlyCostCents int) (*InstanceType, error) {
	if name == "" ||
		cpuCount <= 0 ||
		memoryGB <= 0 ||
		architecture == "" ||
		hourlyCostCents < 0 {
		return nil, ErrInvalidInstanceType
	}

	return &InstanceType{
		name:            name,
		cpuCount:        cpuCount,
		memoryGB:        memoryGB,
		architecture:    architecture,
		hourlyCostCents: hourlyCostCents,
	}, nil
}

func (i *InstanceType) Name() string {
	return i.name
}

func (i *InstanceType) CPUCount() int {
	return i.cpuCount
}

func (i *InstanceType) MemoryGB() int {
	return i.memoryGB
}

func (i *InstanceType) Architecture() string {
	return i.architecture
}

func (i *InstanceType) HourlyCostCents() int {
	return i.hourlyCostCents
}

func (i *InstanceType) CostCents(hours int) int {
	return i.hourlyCostCents * hours
}
