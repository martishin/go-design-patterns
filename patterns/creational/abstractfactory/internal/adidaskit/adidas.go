package adidaskit

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/pkg/sportskit"
)

// AdidasFactory is a concrete factory.
type AdidasFactory struct{}

func New() sportskit.Factory {
	return AdidasFactory{}
}

type adidasShoe struct {
	logo string
	size int
}

func (s adidasShoe) Logo() string {
	return s.logo
}

func (s adidasShoe) Size() int {
	return s.size
}

func (s adidasShoe) String() string {
	return fmt.Sprintf("adidas shoe - logo: %s, size: %d", s.logo, s.size)
}

type adidasShirt struct {
	logo string
	size int
}

func (s adidasShirt) Logo() string {
	return s.logo
}

func (s adidasShirt) Size() int {
	return s.size
}

func (s adidasShirt) String() string {
	return fmt.Sprintf("adidas shirt - logo: %s, size: %d", s.logo, s.size)
}

func (AdidasFactory) Shoe(size int) sportskit.Shoe {
	return adidasShoe{
		logo: "adidas",
		size: size,
	}
}

func (AdidasFactory) Shirt(size int) sportskit.Shirt {
	return adidasShirt{
		logo: "adidas",
		size: size,
	}
}
