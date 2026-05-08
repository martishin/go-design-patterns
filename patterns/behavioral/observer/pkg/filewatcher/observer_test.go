package filewatcher_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

type observerStub struct {
	event filewatcher.Event
}

func (o *observerStub) ID() string {
	return "observer"
}

func (o *observerStub) Update(event filewatcher.Event) {
	o.event = event
}

type subjectStub struct {
	observer filewatcher.Observer
	event    filewatcher.Event
}

func (s *subjectStub) Subscribe(observer filewatcher.Observer) error {
	s.observer = observer
	return nil
}

func (s *subjectStub) Unsubscribe(observer filewatcher.Observer) {
	if s.observer == observer {
		s.observer = nil
	}
}

func (s *subjectStub) Notify(event filewatcher.Event) {
	s.event = event
}

func TestObserver_InterfaceCanBeImplemented(t *testing.T) {
	observer := &observerStub{}
	var typed filewatcher.Observer = observer

	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}
	typed.Update(event)

	if typed.ID() != "observer" {
		t.Fatalf("got observer ID %q, want %q", typed.ID(), "observer")
	}
	if observer.event != event {
		t.Fatalf("got event %v, want %v", observer.event, event)
	}
}

func TestSubject_InterfaceCanBeImplemented(t *testing.T) {
	subject := &subjectStub{}
	observer := &observerStub{}
	var typed filewatcher.Subject = subject

	if err := typed.Subscribe(observer); err != nil {
		t.Fatalf("Subscribe() returned error: %v", err)
	}
	if subject.observer != observer {
		t.Fatalf("got observer %v, want %v", subject.observer, observer)
	}

	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}
	typed.Notify(event)
	if subject.event != event {
		t.Fatalf("got event %v, want %v", subject.event, event)
	}

	typed.Unsubscribe(observer)
	if subject.observer != nil {
		t.Fatalf("got observer %v, want nil", subject.observer)
	}
}
