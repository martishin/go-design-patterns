package main

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/internal/adidaskit"
	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/internal/nikekit"
	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/pkg/sportskit"
)

const (
	shoeSizeEU  = 44
	shirtSizeEU = 52
)

func printKit(brand string, f sportskit.Factory, shoeSize, shirtSize int) {
	fmt.Println("==", brand, "kit ==")
	fmt.Println(f.Shoe(shoeSize))
	fmt.Println(f.Shirt(shirtSize))
}

func main() {
	printKit("Nike", nikekit.New(), shoeSizeEU, shirtSizeEU)
	printKit("Adidas", adidaskit.New(), shoeSizeEU, shirtSizeEU)
}
