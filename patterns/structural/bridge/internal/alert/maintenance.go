package alert

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/pkg/alerts"
)

type MaintenanceAlert struct {
	notifier *alerts.Notifier
	node     string
	reason   string
}

func NewMaintenanceAlert(channel alerts.Channel, node string, reason string) alerts.Alert {
	return &MaintenanceAlert{
		notifier: alerts.NewNotifier(channel),
		node:     node,
		reason:   reason,
	}
}

func (a *MaintenanceAlert) SetChannel(channel alerts.Channel) {
	a.notifier.SetChannel(channel)
}

func (a *MaintenanceAlert) Notify(destination string) error {
	title := fmt.Sprintf("Node %s was sent to maintenance", a.node)
	body := fmt.Sprintf("Because of %s reason", a.reason)

	return a.notifier.Send(destination, title, body)
}
