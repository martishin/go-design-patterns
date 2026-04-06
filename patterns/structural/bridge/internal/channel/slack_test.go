package channel_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/channel"
)

func TestSlackChannel_SendFormatsMessage(t *testing.T) {
	var buf bytes.Buffer
	slack := channel.NewSlackChannel(&buf)

	err := slack.Send("k8s-alerts", "Critical incident in api-server", "OOM. Investigate immediately.")
	if err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	got := buf.String()
	want := "" +
		"Slack channel: k8s-alerts\n" +
		"*Critical incident in api-server*\n" +
		"OOM. Investigate immediately.\n"

	if got != want {
		t.Fatalf("unexpected output:\n--- got ---\n%q\n--- want ---\n%q", got, want)
	}
}
