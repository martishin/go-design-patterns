package subscriber_test

import (
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/internal/subscriber"
	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

func TestBackupService_ImplementsObserver(t *testing.T) {
	var _ filewatcher.Observer = subscriber.NewBackupService("backup")
}

func TestBackupService_UpdateStoresBackupForWriteEvent(t *testing.T) {
	backup := subscriber.NewBackupService("backup")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}

	backup.Update(event)

	want := []string{"backup created for app/config.yaml at revision 1"}
	if !reflect.DeepEqual(backup.Backups(), want) {
		t.Fatalf("got backups %v, want %v", backup.Backups(), want)
	}
}

func TestBackupService_UpdateIgnoresDeleteEvent(t *testing.T) {
	backup := subscriber.NewBackupService("backup")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "delete"}

	backup.Update(event)

	if len(backup.Backups()) != 0 {
		t.Fatalf("got backups %v, want none", backup.Backups())
	}
}

func TestBackupService_BackupsReturnsCopy(t *testing.T) {
	backup := subscriber.NewBackupService("backup")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}
	backup.Update(event)

	backups := backup.Backups()
	backups[0] = "changed"

	if backup.Backups()[0] != "backup created for app/config.yaml at revision 1" {
		t.Fatalf("backups were mutated through returned slice: %v", backup.Backups())
	}
}
