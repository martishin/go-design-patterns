package alerts_test

import (
	"errors"
	"testing"

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

func TestNotifier_SendReturnsErrNoChannel(t *testing.T) {
	notifier := alerts.NewNotifier(nil)

	err := notifier.Send("ops@company.com", "title", "body")
	if !errors.Is(err, alerts.ErrNoChannel) {
		t.Fatalf("got error %v, want %v", err, alerts.ErrNoChannel)
	}
}

func TestNotifier_SendDelegatesToChannel(t *testing.T) {
	spy := &channelSpy{}
	notifier := alerts.NewNotifier(spy)

	err := notifier.Send("ops@company.com", "critical incident", "investigate immediately")
	if err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}
	if spy.calls != 1 {
		t.Fatalf("got %d calls, want 1", spy.calls)
	}
	if spy.gotDestination != "ops@company.com" {
		t.Fatalf("got destination %q, want %q", spy.gotDestination, "ops@company.com")
	}
	if spy.gotTitle != "critical incident" {
		t.Fatalf("got title %q, want %q", spy.gotTitle, "critical incident")
	}
	if spy.gotBody != "investigate immediately" {
		t.Fatalf("got body %q, want %q", spy.gotBody, "investigate immediately")
	}
}
