package alert_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/alert"
	"github.com/martishin/go-design-patterns/patterns/structural/bridge/pkg/alerts"
)

type channelSpy struct {
	gotDestination string
	gotTitle       string
	gotBody        string
	err            error
	calls          int
}

func (s *channelSpy) Send(destination string, title string, body string) error {
	s.gotDestination = destination
	s.gotTitle = title
	s.gotBody = body
	s.calls++

	return s.err
}

func TestIncidentAlert_NotifyBuildsIncidentMessage(t *testing.T) {
	spy := &channelSpy{}
	incident := alert.NewIncidentAlert(spy, "api-server", "OOM")

	err := incident.Notify("k8s-alerts")
	if err != nil {
		t.Fatalf("Notify() returned error: %v", err)
	}
	if spy.calls != 1 {
		t.Fatalf("got %d calls, want 1", spy.calls)
	}
	if spy.gotDestination != "k8s-alerts" {
		t.Fatalf("got destination %q, want %q", spy.gotDestination, "k8s-alerts")
	}
	if spy.gotTitle != "Critical incident in api-server" {
		t.Fatalf("got title %q, want %q", spy.gotTitle, "Critical incident in api-server")
	}
	if spy.gotBody != "OOM. Investigate immediately." {
		t.Fatalf("got body %q, want %q", spy.gotBody, "OOM. Investigate immediately.")
	}
}

func TestIncidentAlert_SetChannelSwitchesImplementation(t *testing.T) {
	first := &channelSpy{}
	second := &channelSpy{}

	incident := alert.NewIncidentAlert(first, "api-server", "OOM")

	if err := incident.Notify("k8s-alerts"); err != nil {
		t.Fatalf("first Notify() returned error: %v", err)
	}

	incident.SetChannel(second)

	if err := incident.Notify("k8s-fallback-alerts"); err != nil {
		t.Fatalf("second Notify() returned error: %v", err)
	}

	if first.calls != 1 {
		t.Fatalf("first channel got %d calls, want 1", first.calls)
	}
	if second.calls != 1 {
		t.Fatalf("second channel got %d calls, want 1", second.calls)
	}
	if second.gotDestination != "k8s-fallback-alerts" {
		t.Fatalf("got destination %q, want %q", second.gotDestination, "k8s-fallback-alerts")
	}
}

var _ alerts.Channel = (*channelSpy)(nil)
