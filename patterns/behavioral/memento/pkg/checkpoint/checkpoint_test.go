package checkpoint_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/pkg/checkpoint"
)

type checkpointStub struct {
	name string
}

func (c checkpointStub) Name() string {
	return c.name
}

func TestCheckpoint_InterfaceCanBeImplemented(t *testing.T) {
	var saved checkpoint.Checkpoint = checkpointStub{name: "before boss fight"}

	if saved.Name() != "before boss fight" {
		t.Fatalf("got checkpoint name %q, want %q", saved.Name(), "before boss fight")
	}
}
