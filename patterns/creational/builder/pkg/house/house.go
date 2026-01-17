package house

import "fmt"

// House is the product.
type House struct {
	WindowType string
	DoorType   string
	NumFloors  int
}

func (h House) String() string {
	return fmt.Sprintf("House with %s windows, %s doors and %d floor(s)", h.WindowType, h.DoorType, h.NumFloors)
}

// Builder is the builder interface.
type Builder interface {
	SetWindowType(string) Builder
	SetDoorType(string) Builder
	SetNumFloors(int) Builder
	Build() House
}
