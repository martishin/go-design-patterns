package deploy_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

type callOrderRecorder struct {
	events []string
}

func (r *callOrderRecorder) add(event string) {
	r.events = append(r.events, event)
}

type validatorSpy struct {
	recorder *callOrderRecorder
	err      error
	calls    int
}

func (s *validatorSpy) Validate(request deploy.DeploymentRequest) error {
	s.calls++
	s.recorder.add("validate")
	return s.err
}

type builderSpy struct {
	recorder *callOrderRecorder
	artifact deploy.Artifact
	err      error
	calls    int
}

func (s *builderSpy) Build(request deploy.DeploymentRequest) (deploy.Artifact, error) {
	s.calls++
	s.recorder.add("build")
	return s.artifact, s.err
}

type artifactStoreSpy struct {
	recorder *callOrderRecorder
	ref      deploy.ArtifactRef
	err      error
	calls    int
}

func (s *artifactStoreSpy) Upload(artifact deploy.Artifact) (deploy.ArtifactRef, error) {
	s.calls++
	s.recorder.add("upload")
	return s.ref, s.err
}

type clusterDeployerSpy struct {
	recorder    *callOrderRecorder
	gotRequest  deploy.DeploymentRequest
	gotArtifact deploy.ArtifactRef
	err         error
	calls       int
}

func (s *clusterDeployerSpy) Deploy(request deploy.DeploymentRequest, artifact deploy.ArtifactRef) error {
	s.calls++
	s.recorder.add("deploy")
	s.gotRequest = request
	s.gotArtifact = artifact
	return s.err
}

type healthCheckerSpy struct {
	recorder       *callOrderRecorder
	gotService     string
	gotEnvironment string
	err            error
	calls          int
}

func (s *healthCheckerSpy) Check(service string, environment string) error {
	s.calls++
	s.recorder.add("healthcheck")
	s.gotService = service
	s.gotEnvironment = environment
	return s.err
}

type notifierSpy struct {
	recorder    *callOrderRecorder
	gotRequest  deploy.DeploymentRequest
	gotArtifact deploy.ArtifactRef
	err         error
	calls       int
}

func (s *notifierSpy) NotifySuccess(request deploy.DeploymentRequest, artifact deploy.ArtifactRef) error {
	s.calls++
	s.recorder.add("notify")
	s.gotRequest = request
	s.gotArtifact = artifact
	return s.err
}

func TestNewDeploymentFacade_ReturnsErrNilValidator(t *testing.T) {
	_, err := deploy.NewDeploymentFacade(
		nil,
		&builderSpy{},
		&artifactStoreSpy{},
		&clusterDeployerSpy{},
		&healthCheckerSpy{},
		&notifierSpy{},
	)
	if !errors.Is(err, deploy.ErrNilValidator) {
		t.Fatalf("got error %v, want %v", err, deploy.ErrNilValidator)
	}
}

func TestDeploymentFacade_DeployCoordinatesSubsystemsInOrder(t *testing.T) {
	recorder := &callOrderRecorder{}
	request := deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "production",
	}
	artifact := deploy.Artifact{
		Service: request.Service,
		Version: request.Version,
		Image:   "payments-api:v1.4.2",
	}
	ref := deploy.ArtifactRef{URI: "registry.acme.io/payments-api:v1.4.2"}

	validatorSpy := &validatorSpy{recorder: recorder}
	builderSpy := &builderSpy{recorder: recorder, artifact: artifact}
	artifactStoreSpy := &artifactStoreSpy{recorder: recorder, ref: ref}
	clusterSpy := &clusterDeployerSpy{recorder: recorder}
	healthCheckerSpy := &healthCheckerSpy{recorder: recorder}
	notifierSpy := &notifierSpy{recorder: recorder}

	facade, err := deploy.NewDeploymentFacade(
		validatorSpy,
		builderSpy,
		artifactStoreSpy,
		clusterSpy,
		healthCheckerSpy,
		notifierSpy,
	)
	if err != nil {
		t.Fatalf("NewDeploymentFacade() returned error: %v", err)
	}

	if err := facade.Deploy(request); err != nil {
		t.Fatalf("Deploy() returned error: %v", err)
	}

	gotOrder := recorder.events
	wantOrder := []string{"validate", "build", "upload", "deploy", "healthcheck", "notify"}
	if len(gotOrder) != len(wantOrder) {
		t.Fatalf("got call order %v, want %v", gotOrder, wantOrder)
	}
	for i := range wantOrder {
		if gotOrder[i] != wantOrder[i] {
			t.Fatalf("got call order %v, want %v", gotOrder, wantOrder)
		}
	}

	if clusterSpy.gotRequest != request {
		t.Fatalf("got deploy request %+v, want %+v", clusterSpy.gotRequest, request)
	}
	if clusterSpy.gotArtifact != ref {
		t.Fatalf("got artifact ref %+v, want %+v", clusterSpy.gotArtifact, ref)
	}
	if healthCheckerSpy.gotService != request.Service {
		t.Fatalf("got service %q, want %q", healthCheckerSpy.gotService, request.Service)
	}
	if healthCheckerSpy.gotEnvironment != request.Environment {
		t.Fatalf("got environment %q, want %q", healthCheckerSpy.gotEnvironment, request.Environment)
	}
	if notifierSpy.gotRequest != request {
		t.Fatalf("got notification request %+v, want %+v", notifierSpy.gotRequest, request)
	}
	if notifierSpy.gotArtifact != ref {
		t.Fatalf("got notification artifact %+v, want %+v", notifierSpy.gotArtifact, ref)
	}
}

func TestDeploymentFacade_DeployStopsWhenValidationFails(t *testing.T) {
	recorder := &callOrderRecorder{}
	wantErr := errors.New("validation failed")
	validatorSpy := &validatorSpy{recorder: recorder, err: wantErr}
	builderSpy := &builderSpy{recorder: recorder}
	artifactStoreSpy := &artifactStoreSpy{recorder: recorder}
	clusterSpy := &clusterDeployerSpy{recorder: recorder}
	healthCheckerSpy := &healthCheckerSpy{recorder: recorder}
	notifierSpy := &notifierSpy{recorder: recorder}

	facade, err := deploy.NewDeploymentFacade(
		validatorSpy,
		builderSpy,
		artifactStoreSpy,
		clusterSpy,
		healthCheckerSpy,
		notifierSpy,
	)
	if err != nil {
		t.Fatalf("NewDeploymentFacade() returned error: %v", err)
	}

	err = facade.Deploy(deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "production",
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
	if builderSpy.calls != 0 || artifactStoreSpy.calls != 0 || clusterSpy.calls != 0 || healthCheckerSpy.calls != 0 || notifierSpy.calls != 0 {
		t.Fatalf(
			"expected no downstream calls, got builder=%d upload=%d deploy=%d healthcheck=%d notify=%d",
			builderSpy.calls,
			artifactStoreSpy.calls,
			clusterSpy.calls,
			healthCheckerSpy.calls,
			notifierSpy.calls,
		)
	}
}
