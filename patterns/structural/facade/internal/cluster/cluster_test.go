package cluster_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/cluster"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func TestCluster_DeployReturnsErrEmptyArtifactURI(t *testing.T) {
	clusterDeployer := cluster.NewCluster("production-eu")

	err := clusterDeployer.Deploy(
		deploy.DeploymentRequest{
			Service:     "payments-api",
			Version:     "v1.4.2",
			Environment: "production",
		},
		deploy.ArtifactRef{},
	)
	if !errors.Is(err, cluster.ErrEmptyArtifactURI) {
		t.Fatalf("got error %v, want %v", err, cluster.ErrEmptyArtifactURI)
	}
}

func TestCluster_DeployAcceptsArtifactRef(t *testing.T) {
	clusterDeployer := cluster.NewCluster("production-eu")

	err := clusterDeployer.Deploy(
		deploy.DeploymentRequest{
			Service:     "payments-api",
			Version:     "v1.4.2",
			Environment: "production",
		},
		deploy.ArtifactRef{URI: "registry.acme.io/payments-api:v1.4.2"},
	)
	if err != nil {
		t.Fatalf("Deploy() returned error: %v", err)
	}
}
