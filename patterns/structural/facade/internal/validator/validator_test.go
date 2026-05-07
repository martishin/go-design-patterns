package validator_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/validator"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func TestReleaseValidator_ValidateReturnsErrEmptyService(t *testing.T) {
	releaseValidator := validator.NewReleaseValidator()

	err := releaseValidator.Validate(deploy.DeploymentRequest{
		Service:     "",
		Version:     "v1.4.2",
		Environment: "production",
	})
	if !errors.Is(err, validator.ErrEmptyService) {
		t.Fatalf("got error %v, want %v", err, validator.ErrEmptyService)
	}
}

func TestReleaseValidator_ValidateReturnsErrUnsupportedEnvironment(t *testing.T) {
	releaseValidator := validator.NewReleaseValidator()

	err := releaseValidator.Validate(deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "qa",
	})
	if !errors.Is(err, validator.ErrUnsupportedEnvironment) {
		t.Fatalf("got error %v, want %v", err, validator.ErrUnsupportedEnvironment)
	}
}

func TestReleaseValidator_ValidateAcceptsSupportedRequest(t *testing.T) {
	releaseValidator := validator.NewReleaseValidator()

	err := releaseValidator.Validate(deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "production",
	})
	if err != nil {
		t.Fatalf("Validate() returned error: %v", err)
	}
}
