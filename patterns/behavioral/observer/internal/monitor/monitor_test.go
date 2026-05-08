package monitor_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/internal/monitor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

type observerSpy struct {
	id     string
	events []filewatcher.Event
}

func (o *observerSpy) ID() string {
	return o.id
}

func (o *observerSpy) Update(event filewatcher.Event) {
	o.events = append(o.events, event)
}

func TestFileMonitor_ImplementsSubject(t *testing.T) {
	var _ filewatcher.Subject = monitor.NewFileMonitor()
}

func TestFileMonitor_SubscribeReturnsErrorForNilObserver(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()

	err := fileMonitor.Subscribe(nil)

	if !errors.Is(err, monitor.ErrNilObserver) {
		t.Fatalf("got error %v, want %v", err, monitor.ErrNilObserver)
	}
}

func TestFileMonitor_SubscribeIgnoresDuplicateObserver(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()
	observer := &observerSpy{id: "indexer"}

	if err := fileMonitor.Subscribe(observer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}
	if err := fileMonitor.Subscribe(observer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}

	if fileMonitor.SubscriberCount() != 1 {
		t.Fatalf("got subscriber count %d, want %d", fileMonitor.SubscriberCount(), 1)
	}
}

func TestFileMonitor_UnsubscribeRemovesObserver(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()
	indexer := &observerSpy{id: "indexer"}
	backup := &observerSpy{id: "backup"}

	if err := fileMonitor.Subscribe(indexer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}
	if err := fileMonitor.Subscribe(backup); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}

	fileMonitor.Unsubscribe(indexer)
	if err := fileMonitor.Write("app/config.yaml"); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	if len(indexer.events) != 0 {
		t.Fatalf("got indexer events %v, want none", indexer.events)
	}

	wantBackupEvents := []filewatcher.Event{
		{Path: "app/config.yaml", Revision: 1, Operation: "write"},
	}
	if !reflect.DeepEqual(backup.events, wantBackupEvents) {
		t.Fatalf("got backup events %v, want %v", backup.events, wantBackupEvents)
	}
}

func TestFileMonitor_WriteNotifiesObservers(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()
	indexer := &observerSpy{id: "indexer"}
	backup := &observerSpy{id: "backup"}

	if err := fileMonitor.Subscribe(indexer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}
	if err := fileMonitor.Subscribe(backup); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}

	if err := fileMonitor.Write("app/config.yaml"); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	wantEvents := []filewatcher.Event{
		{Path: "app/config.yaml", Revision: 1, Operation: "write"},
	}
	if !reflect.DeepEqual(indexer.events, wantEvents) {
		t.Fatalf("got indexer events %v, want %v", indexer.events, wantEvents)
	}
	if !reflect.DeepEqual(backup.events, wantEvents) {
		t.Fatalf("got backup events %v, want %v", backup.events, wantEvents)
	}
}

func TestFileMonitor_DeleteNotifiesObservers(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()
	indexer := &observerSpy{id: "indexer"}

	if err := fileMonitor.Subscribe(indexer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}

	if err := fileMonitor.Delete("app/config.yaml"); err != nil {
		t.Fatalf("Delete() returned error: %v", err)
	}

	wantEvents := []filewatcher.Event{
		{Path: "app/config.yaml", Revision: 1, Operation: "delete"},
	}
	if !reflect.DeepEqual(indexer.events, wantEvents) {
		t.Fatalf("got indexer events %v, want %v", indexer.events, wantEvents)
	}
}

func TestFileMonitor_ReturnsErrorForEmptyPath(t *testing.T) {
	fileMonitor := monitor.NewFileMonitor()

	err := fileMonitor.Write("")

	if !errors.Is(err, monitor.ErrEmptyPath) {
		t.Fatalf("got error %v, want %v", err, monitor.ErrEmptyPath)
	}
}
