package main

import (
	"os"

	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/alert"
	"github.com/martishin/go-design-patterns/patterns/structural/bridge/internal/channel"
)

func main() {
	slackChannel := channel.NewSlackChannel(os.Stdout)
	emailChannel := channel.NewEmailChannel(os.Stdout)

	incidentAlert := alert.NewIncidentAlert(slackChannel, "api-server", "OOM")
	nodeAlert := alert.NewMaintenanceAlert(emailChannel, "node-1234", "switch update")

	_ = incidentAlert.Notify("k8s-alerts")
	_ = nodeAlert.Notify("nodes-alerts@acme.com")

	nodeAlert.SetChannel(slackChannel)

	_ = nodeAlert.Notify("nodes-alerts")
}
