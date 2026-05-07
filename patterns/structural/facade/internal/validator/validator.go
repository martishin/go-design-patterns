package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

var (
	ErrEmptyService           = errors.New("validator: service is required")
	ErrEmptyVersion           = errors.New("validator: version is required")
	ErrEmptyEnvironment       = errors.New("validator: environment is required")
	ErrUnsupportedEnvironment = errors.New("validator: unsupported environment")
)

type ReleaseValidator struct{}

func NewReleaseValidator() deploy.Validator {
	return &ReleaseValidator{}
}

func (v *ReleaseValidator) Validate(request deploy.DeploymentRequest) error {
	switch {
	case strings.TrimSpace(request.Service) == "":
		return ErrEmptyService
	case strings.TrimSpace(request.Version) == "":
		return ErrEmptyVersion
	case strings.TrimSpace(request.Environment) == "":
		return ErrEmptyEnvironment
	}

	if !isSupportedEnvironment(request.Environment) {
		return fmt.Errorf("%w: %s", ErrUnsupportedEnvironment, request.Environment)
	}

	fmt.Printf(
		"Validated deployment request for service %s version %s in %s\n",
		request.Service,
		request.Version,
		request.Environment,
	)

	return nil
}

func isSupportedEnvironment(environment string) bool {
	switch environment {
	case "development", "staging", "production":
		return true
	default:
		return false
	}
}
