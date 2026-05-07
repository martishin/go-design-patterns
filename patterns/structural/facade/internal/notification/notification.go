package notification

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

type DeploymentNotifier struct {
	out io.Writer
}

func NewDeploymentNotifier(out io.Writer) deploy.Notifier {
	return &DeploymentNotifier{out: out}
}

func (n *DeploymentNotifier) NotifySuccess(request deploy.DeploymentRequest, artifact deploy.ArtifactRef) error {
	_, err := fmt.Fprintf(
		n.out,
		"Deployment succeeded for %s version %s in %s with artifact %s\n",
		request.Service,
		request.Version,
		request.Environment,
		artifact.URI,
	)

	return err
}
