package channel

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/pkg/alerts"
)

type EmailChannel struct {
	out io.Writer
}

func (e EmailChannel) Send(destination string, title string, body string) error {
	_, err := fmt.Fprintf(
		e.out,
		"Email to: %s\nSubject: %s\n%s\n",
		destination,
		title,
		body,
	)

	return err
}

func NewEmailChannel(out io.Writer) alerts.Channel {
	return &EmailChannel{out: out}
}
