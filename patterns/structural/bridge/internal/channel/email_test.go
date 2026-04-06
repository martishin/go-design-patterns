package channel_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/channel"
)

func TestEmailChannel_SendFormatsMessage(t *testing.T) {
	var buf bytes.Buffer
	email := channel.NewEmailChannel(&buf)

	err := email.Send("ops@company.com", "Critical incident in api-server", "OOM. Investigate immediately.")
	if err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	got := buf.String()
	want := "" +
		"Email to: ops@company.com\n" +
		"Subject: Critical incident in api-server\n" +
		"OOM. Investigate immediately.\n"

	if got != want {
		t.Fatalf("unexpected output:\n--- got ---\n%q\n--- want ---\n%q", got, want)
	}
}
