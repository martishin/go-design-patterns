package builders_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/builder/internal/builders"
)

func TestNormalBuilder(t *testing.T) {
	b := builders.NewNormalBuilder()
	h := b.SetWindowType("Wooden").SetDoorType("Wooden").SetNumFloors(2).Build()

	if h.WindowType != "Wooden" {
		t.Errorf("Expected Wooden Window, got %s", h.WindowType)
	}
	if h.DoorType != "Wooden" {
		t.Errorf("Expected Wooden Door, got %s", h.DoorType)
	}
	if h.NumFloors != 2 {
		t.Errorf("Expected 2 floors, got %d", h.NumFloors)
	}
}

func TestIglooBuilder(t *testing.T) {
	b := builders.NewIglooBuilder()
	h := b.Build()

	if h.WindowType != "Snow" {
		t.Errorf("Expected Snow Window, got %s", h.WindowType)
	}
	if h.DoorType != "Snow" {
		t.Errorf("Expected Snow Door, got %s", h.DoorType)
	}
	if h.NumFloors != 1 {
		t.Errorf("Expected 1 numFloors, got %d", h.NumFloors)
	}
}
