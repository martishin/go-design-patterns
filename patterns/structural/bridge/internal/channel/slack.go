package channel

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/pkg/alerts"
)

type SlackChannel struct {
	out io.Writer
}

func NewSlackChannel(out io.Writer) alerts.Channel {
	return &SlackChannel{
		out: out,
	}
}

func (s SlackChannel) Send(destination string, title string, body string) error {
	_, err := fmt.Fprintf(
		s.out,
		"Slack channel: %s\n*%s*\n%s\n",
		destination,
		title,
		body,
	)

	return err
}
