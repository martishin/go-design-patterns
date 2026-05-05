package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/structural/composite/internal/infrastructure"
	"github.com/martishin/go-design-patterns/patterns/structural/composite/pkg/cost"
)

func main() {
	apiServer, err := infrastructure.NewComputeInstance("api-server", 4)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := infrastructure.NewManagedDatabase("postgres", 200)
	if err != nil {
		log.Fatal(err)
	}

	production, err := cost.NewCostGroup("production", apiServer, postgres)
	if err != nil {
		log.Fatal(err)
	}

	stagingAPI, err := infrastructure.NewComputeInstance("staging-api", 2)
	if err != nil {
		log.Fatal(err)
	}

	stagingPostgres, err := infrastructure.NewManagedDatabase("staging-postgres", 50)
	if err != nil {
		log.Fatal(err)
	}

	staging, err := cost.NewCostGroup("staging", stagingAPI, stagingPostgres)
	if err != nil {
		log.Fatal(err)
	}

	allEnvironments, err := cost.NewCostGroup("all-environments", production, staging)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Production monthly cost: %s\n", formatCost(production.MonthlyCostCents()))
	fmt.Printf("All environments monthly cost: %s\n", formatCost(allEnvironments.MonthlyCostCents()))
}

func formatCost(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}
