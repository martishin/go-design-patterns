package builders

import "github.com/martishin/go-design-patterns/patterns/creational/builder/pkg/house"

type normalBuilder struct {
	windowType string
	doorType   string
	numFloors  int
}

func NewNormalBuilder() house.Builder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetWindowType(windowType string) house.Builder {
	b.windowType = windowType
	return b
}

func (b *normalBuilder) SetDoorType(doorType string) house.Builder {
	b.doorType = doorType
	return b
}

func (b *normalBuilder) SetNumFloors(numFloors int) house.Builder {
	b.numFloors = numFloors
	return b
}

func (b *normalBuilder) Build() house.House {
	return house.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		NumFloors:  b.numFloors,
	}
}
