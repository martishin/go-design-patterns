package cluster

import (
	"errors"
	"fmt"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

var (
	ErrEmptyClusterName = errors.New("cluster: name is required")
	ErrEmptyArtifactURI = errors.New("cluster: artifact URI is required")
)

type Cluster struct {
	name string
}

func NewCluster(name string) deploy.ClusterDeployer {
	return &Cluster{name: name}
}

func (c *Cluster) Deploy(request deploy.DeploymentRequest, artifact deploy.ArtifactRef) error {
	switch {
	case strings.TrimSpace(c.name) == "":
		return ErrEmptyClusterName
	case strings.TrimSpace(artifact.URI) == "":
		return ErrEmptyArtifactURI
	}

	fmt.Printf(
		"Deploying %s to %s on cluster %s\n",
		artifact.URI,
		request.Environment,
		c.name,
	)

	return nil
}
