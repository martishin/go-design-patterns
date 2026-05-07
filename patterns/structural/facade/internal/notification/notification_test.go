package notification_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/notification"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func TestDeploymentNotifier_NotifySuccessFormatsMessage(t *testing.T) {
	var buf bytes.Buffer
	notifier := notification.NewDeploymentNotifier(&buf)

	err := notifier.NotifySuccess(
		deploy.DeploymentRequest{
			Service:     "payments-api",
			Version:     "v1.4.2",
			Environment: "production",
		},
		deploy.ArtifactRef{URI: "registry.acme.io/payments-api:v1.4.2"},
	)
	if err != nil {
		t.Fatalf("NotifySuccess() returned error: %v", err)
	}

	got := buf.String()
	want := "Deployment succeeded for payments-api version v1.4.2 in production with artifact registry.acme.io/payments-api:v1.4.2\n"
	if got != want {
		t.Fatalf("got message %q, want %q", got, want)
	}
}
