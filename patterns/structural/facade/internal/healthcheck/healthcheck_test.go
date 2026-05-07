package healthcheck_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/healthcheck"
)

func TestHTTPHealthChecker_CheckReturnsErrEmptyServiceName(t *testing.T) {
	checker := healthcheck.NewHTTPHealthChecker()

	err := checker.Check("", "production")
	if !errors.Is(err, healthcheck.ErrEmptyServiceName) {
		t.Fatalf("got error %v, want %v", err, healthcheck.ErrEmptyServiceName)
	}
}

func TestHTTPHealthChecker_CheckAcceptsServiceAndEnvironment(t *testing.T) {
	checker := healthcheck.NewHTTPHealthChecker()

	err := checker.Check("payments-api", "production")
	if err != nil {
		t.Fatalf("Check() returned error: %v", err)
	}
}
