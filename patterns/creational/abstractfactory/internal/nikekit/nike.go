package nikekit

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/pkg/sportskit"
)

// NikeFactory is a concrete factory.
type NikeFactory struct{}

func New() sportskit.Factory {
	return NikeFactory{}
}

type nikeShoe struct {
	logo string
	size int
}

func (s nikeShoe) Logo() string {
	return s.logo
}

func (s nikeShoe) Size() int {
	return s.size
}

func (s nikeShoe) String() string {
	return fmt.Sprintf("nike shoe - logo: %s, size: %d", s.logo, s.size)
}

type nikeShirt struct {
	logo string
	size int
}

func (s nikeShirt) Logo() string {
	return s.logo
}

func (s nikeShirt) Size() int {
	return s.size
}

func (s nikeShirt) String() string {
	return fmt.Sprintf("nike shirt - logo: %s, size: %d", s.logo, s.size)
}

func (NikeFactory) Shoe(size int) sportskit.Shoe {
	return nikeShoe{
		logo: "nike",
		size: size,
	}
}

func (NikeFactory) Shirt(size int) sportskit.Shirt {
	return nikeShirt{
		logo: "nike",
		size: size,
	}
}
