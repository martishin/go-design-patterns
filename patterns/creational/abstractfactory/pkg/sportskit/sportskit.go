package sportskit

// Shoe and Shirt are the abstract products.
type Shoe interface {
	Logo() string
	Size() int
}

type Shirt interface {
	Logo() string
	Size() int
}

// Factory is the abstract factory that creates brand-matched products.
type Factory interface {
	Shoe(size int) Shoe
	Shirt(size int) Shirt
}
