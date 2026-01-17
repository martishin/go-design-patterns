package house

// Architect is responsible for managing the building process.
type Architect struct {
	builder Builder
}

func NewArchitect(b Builder) *Architect {
	return &Architect{
		builder: b,
	}
}

func (d *Architect) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Architect) SetHouseParams(windowType string, doorType string, numFloors int) {
	d.builder.
		SetWindowType(windowType).
		SetDoorType(doorType).
		SetNumFloors(numFloors)
}

func (d *Architect) BuildHouse() House {
	return d.builder.Build()
}
