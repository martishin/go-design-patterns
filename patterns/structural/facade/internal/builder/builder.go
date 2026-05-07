package builder

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

type ImageBuilder struct{}

func NewImageBuilder() deploy.Builder {
	return &ImageBuilder{}
}

func (b *ImageBuilder) Build(request deploy.DeploymentRequest) (deploy.Artifact, error) {
	artifact := deploy.Artifact{
		Service: request.Service,
		Version: request.Version,
		Image:   fmt.Sprintf("%s:%s", request.Service, request.Version),
	}

	fmt.Printf("Built artifact %s\n", artifact.Image)

	return artifact, nil
}
