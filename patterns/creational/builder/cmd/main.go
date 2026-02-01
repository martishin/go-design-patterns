package main

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/builder/internal/builders"
	"github.com/martishin/go-design-patterns/patterns/creational/builder/pkg/house"
)

const numOfFloors = 2

func main() {
	normalBuilder := builders.NewNormalBuilder()
	iglooBuilder := builders.NewIglooBuilder()

	architect := house.NewArchitect(normalBuilder)
	architect.SetHouseParams("Wooden", "Wooden", numOfFloors)

	normalHouse := architect.BuildHouse()
	fmt.Println(normalHouse)

	architect.SetBuilder(iglooBuilder)

	iglooHouse := architect.BuildHouse()
	fmt.Println(iglooHouse)
}
