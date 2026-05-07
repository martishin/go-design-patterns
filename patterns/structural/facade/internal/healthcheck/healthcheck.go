package healthcheck

import (
	"errors"
	"fmt"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

var (
	ErrEmptyServiceName     = errors.New("healthcheck: service is required")
	ErrEmptyEnvironmentName = errors.New("healthcheck: environment is required")
)

type HTTPHealthChecker struct{}

func NewHTTPHealthChecker() deploy.HealthChecker {
	return &HTTPHealthChecker{}
}

func (h *HTTPHealthChecker) Check(service string, environment string) error {
	switch {
	case strings.TrimSpace(service) == "":
		return ErrEmptyServiceName
	case strings.TrimSpace(environment) == "":
		return ErrEmptyEnvironmentName
	}

	fmt.Printf("Health check passed for %s in %s\n", service, environment)

	return nil
}
