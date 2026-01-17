package builders

import "github.com/martishin/go-design-patterns/patterns/creational/builder/pkg/house"

type iglooBuilder struct {
	windowType string
	doorType   string
	numFloors  int
}

func NewIglooBuilder() house.Builder {
	return &iglooBuilder{
		windowType: "Snow",
		doorType:   "Snow",
		numFloors:  1,
	}
}

func (b *iglooBuilder) SetWindowType(windowType string) house.Builder {
	b.windowType = windowType
	return b
}

func (b *iglooBuilder) SetDoorType(doorType string) house.Builder {
	b.doorType = doorType
	return b
}

func (b *iglooBuilder) SetNumFloors(numFloors int) house.Builder {
	b.numFloors = numFloors
	return b
}

func (b *iglooBuilder) Build() house.House {
	return house.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		NumFloors:  b.numFloors,
	}
}
