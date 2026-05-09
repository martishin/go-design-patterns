package review_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/pkg/review"
)

type stateStub struct {
	name string
}

func (s stateStub) Name() string {
	return s.name
}

func (s stateStub) Open() error {
	return nil
}

func (s stateStub) Approve() error {
	return nil
}

func (s stateStub) RequestChanges() error {
	return nil
}

func (s stateStub) Merge() error {
	return nil
}

func TestState_InterfaceCanBeImplemented(t *testing.T) {
	var state review.State = stateStub{name: "draft"}

	if state.Name() != "draft" {
		t.Fatalf("got state name %q, want %q", state.Name(), "draft")
	}
}
