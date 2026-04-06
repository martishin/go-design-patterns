package alert

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/pkg/alerts"
)

type IncidentAlert struct {
	notifier *alerts.Notifier
	service  string
	issue    string
}

func NewIncidentAlert(channel alerts.Channel, service string, issue string) alerts.Alert {
	return &IncidentAlert{
		notifier: alerts.NewNotifier(channel),
		service:  service,
		issue:    issue,
	}
}

func (a *IncidentAlert) SetChannel(channel alerts.Channel) {
	a.notifier.SetChannel(channel)
}

func (a *IncidentAlert) Notify(destination string) error {
	title := fmt.Sprintf("Critical incident in %s", a.service)
	body := fmt.Sprintf("%s. Investigate immediately.", a.issue)

	return a.notifier.Send(destination, title, body)
}
