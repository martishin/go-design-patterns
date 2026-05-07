package deploy

import (
	"errors"
	"reflect"
)

var (
	ErrNilValidator       = errors.New("deploy: nil validator")
	ErrNilBuilder         = errors.New("deploy: nil builder")
	ErrNilArtifactStore   = errors.New("deploy: nil artifact store")
	ErrNilClusterDeployer = errors.New("deploy: nil cluster deployer")
	ErrNilHealthChecker   = errors.New("deploy: nil health checker")
	ErrNilNotifier        = errors.New("deploy: nil notifier")
)

type DeploymentRequest struct {
	Service     string
	Version     string
	Environment string
}

type Artifact struct {
	Service string
	Version string
	Image   string
}

type ArtifactRef struct {
	URI string
}

type Validator interface {
	Validate(request DeploymentRequest) error
}

type Builder interface {
	Build(request DeploymentRequest) (Artifact, error)
}

type ArtifactStore interface {
	Upload(artifact Artifact) (ArtifactRef, error)
}

type ClusterDeployer interface {
	Deploy(request DeploymentRequest, artifact ArtifactRef) error
}

type HealthChecker interface {
	Check(service string, environment string) error
}

type Notifier interface {
	NotifySuccess(request DeploymentRequest, artifact ArtifactRef) error
}

type DeploymentFacade struct {
	validator       Validator
	builder         Builder
	artifactStore   ArtifactStore
	clusterDeployer ClusterDeployer
	healthChecker   HealthChecker
	notifier        Notifier
}

func NewDeploymentFacade(
	validator Validator,
	builder Builder,
	artifactStore ArtifactStore,
	clusterDeployer ClusterDeployer,
	healthChecker HealthChecker,
	notifier Notifier,
) (*DeploymentFacade, error) {
	switch {
	case isNilDependency(validator):
		return nil, ErrNilValidator
	case isNilDependency(builder):
		return nil, ErrNilBuilder
	case isNilDependency(artifactStore):
		return nil, ErrNilArtifactStore
	case isNilDependency(clusterDeployer):
		return nil, ErrNilClusterDeployer
	case isNilDependency(healthChecker):
		return nil, ErrNilHealthChecker
	case isNilDependency(notifier):
		return nil, ErrNilNotifier
	}

	return &DeploymentFacade{
		validator:       validator,
		builder:         builder,
		artifactStore:   artifactStore,
		clusterDeployer: clusterDeployer,
		healthChecker:   healthChecker,
		notifier:        notifier,
	}, nil
}

func (f *DeploymentFacade) Deploy(request DeploymentRequest) error {
	if err := f.validator.Validate(request); err != nil {
		return err
	}

	artifact, err := f.builder.Build(request)
	if err != nil {
		return err
	}

	artifactRef, err := f.artifactStore.Upload(artifact)
	if err != nil {
		return err
	}

	if err := f.clusterDeployer.Deploy(request, artifactRef); err != nil {
		return err
	}

	if err := f.healthChecker.Check(request.Service, request.Environment); err != nil {
		return err
	}

	return f.notifier.NotifySuccess(request, artifactRef)
}

func isNilDependency(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
