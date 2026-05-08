package command_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type commandStub struct {
	name    string
	changed bool
	err     error
	undone  bool
}

func (c *commandStub) Name() string {
	return c.name
}

func (c *commandStub) Description() string {
	return "test description"
}

func (c *commandStub) Execute() (bool, error) {
	return c.changed, c.err
}

func (c *commandStub) Undo() {
	c.undone = true
}

func TestCommand_InterfaceCanBeImplemented(t *testing.T) {
	var cmd command.Command = &commandStub{name: "test command", changed: true}

	if cmd.Name() != "test command" {
		t.Fatalf("got name %q, want %q", cmd.Name(), "test command")
	}
	if cmd.Description() != "test description" {
		t.Fatalf("got description %q, want %q", cmd.Description(), "test description")
	}

	changed, err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	if !changed {
		t.Fatal("got changed=false, want true")
	}
}
