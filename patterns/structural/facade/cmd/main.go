package main

import (
	"log"
	"os"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/builder"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/cluster"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/healthcheck"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/notification"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/registry"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/validator"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func main() {
	deploymentFacade, err := deploy.NewDeploymentFacade(
		validator.NewReleaseValidator(),
		builder.NewImageBuilder(),
		registry.NewRegistry("registry.acme.io"),
		cluster.NewCluster("production-eu"),
		healthcheck.NewHTTPHealthChecker(),
		notification.NewDeploymentNotifier(os.Stdout),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = deploymentFacade.Deploy(deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "production",
	})
	if err != nil {
		log.Fatal(err)
	}
}
