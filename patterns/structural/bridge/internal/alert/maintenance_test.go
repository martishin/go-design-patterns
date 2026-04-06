package alert_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/alert"
)

func TestMaintenanceAlert_NotifyBuildsMaintenanceMessage(t *testing.T) {
	spy := &channelSpy{}
	maintenance := alert.NewMaintenanceAlert(spy, "node-1234", "switch update")

	err := maintenance.Notify("nodes-alerts@acme.com")
	if err != nil {
		t.Fatalf("Notify() returned error: %v", err)
	}
	if spy.calls != 1 {
		t.Fatalf("got %d calls, want 1", spy.calls)
	}
	if spy.gotDestination != "nodes-alerts@acme.com" {
		t.Fatalf("got destination %q, want %q", spy.gotDestination, "nodes-alerts@acme.com")
	}
	if spy.gotTitle != "Node node-1234 was sent to maintenance" {
		t.Fatalf("got title %q, want %q", spy.gotTitle, "Node node-1234 was sent to maintenance")
	}
	if spy.gotBody != "Because of switch update reason" {
		t.Fatalf("got body %q, want %q", spy.gotBody, "Because of switch update reason")
	}
}
