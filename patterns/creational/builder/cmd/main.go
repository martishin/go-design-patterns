package main

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/builder/internal/builders"
	"github.com/martishin/go-design-patterns/patterns/creational/builder/pkg/house"
)

func main() {
	normalBuilder := builders.NewNormalBuilder()
	iglooBuilder := builders.NewIglooBuilder()

	architect := house.NewArchitect(normalBuilder)
	architect.SetHouseParams("Wooden", "Wooden", 2)

	normalHouse := architect.BuildHouse()
	fmt.Println(normalHouse)

	architect.SetBuilder(iglooBuilder)

	iglooHouse := architect.BuildHouse()
	fmt.Println(iglooHouse)
}
