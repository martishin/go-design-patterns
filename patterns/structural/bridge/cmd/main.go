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

// What main() should show
//
// - IncidentAlert + SlackChannel
// - MaintenanceAlert + EmailChannel
// - then switch one alert at runtime with SetChannel(...) and send again

// Tests To Add
//
// - [incident_test.go](/Volumes/SSD/development/code/go-design-patterns/patterns/structural/bridge/internal/incidentalert/incident_test.go)
// - TestIncidentAlert_NotifyBuildsIncidentMessage
// - TestIncidentAlert_SetChannelSwitchesImplementation
// - [maintenance_test.go](/Volumes/SSD/development/code/go-design-patterns/patterns/structural/bridge/internal/maintenancealert/maintenance_test.go)
// - TestMaintenanceAlert_NotifyBuildsMaintenanceMessage
// - [email_test.go](/Volumes/SSD/development/code/go-design-patterns/patterns/structural/bridge/internal/emailchannel/email_test.go)
// - TestEmailChannel_SendFormatsMessage
// - [slack_test.go](/Volumes/SSD/development/code/go-design-patterns/patterns/structural/bridge/internal/slackchannel/slack_test.go)
// - TestSlackChannel_SendFormatsMessage
