package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/internal/fleet"
	"github.com/martishin/go-design-patterns/patterns/structural/flyweight/pkg/cloud"
)

func main() {
	catalog, err := cloud.NewDefaultInstanceTypeCatalog()
	if err != nil {
		log.Fatal(err)
	}

	productionFleet, err := fleet.NewFleet(catalog)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := productionFleet.Launch("api-1", "eu-west-1", "production", "standard-large", 720); err != nil {
		log.Fatal(err)
	}
	if _, err := productionFleet.Launch("api-2", "eu-west-1", "production", "standard-large", 720); err != nil {
		log.Fatal(err)
	}
	if _, err := productionFleet.Launch("worker-1", "eu-west-1", "production", "arm-medium", 360); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Instances: %d\n", len(productionFleet.Instances()))
	fmt.Printf("Shared instance types: %d\n", productionFleet.SharedInstanceTypeCount())
	fmt.Printf("Monthly cost: %s\n", formatCost(productionFleet.TotalMonthlyCostCents()))
}

func formatCost(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}
