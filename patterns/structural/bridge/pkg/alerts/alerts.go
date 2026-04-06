package alerts

import "errors"

type Channel interface {
	Send(destination string, title string, body string) error
}

type Alert interface {
	Notify(destination string) error
	SetChannel(channel Channel)
}

var ErrNoChannel = errors.New("alerts: no channel configured")

type Notifier struct {
	channel Channel
}

func NewNotifier(channel Channel) *Notifier {
	return &Notifier{channel: channel}
}

func (n *Notifier) SetChannel(channel Channel) {
	n.channel = channel
}

func (n *Notifier) Send(destination string, title string, body string) error {
	if n == nil || n.channel == nil {
		return ErrNoChannel
	}

	return n.channel.Send(destination, title, body)
}
