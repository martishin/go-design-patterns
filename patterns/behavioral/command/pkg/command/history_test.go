package command_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

func TestHistory_PopReturnsNilWhenEmpty(t *testing.T) {
	history := command.NewHistory()

	got := history.Pop()
	if got != nil {
		t.Fatalf("got command %v, want nil", got)
	}
}

func TestHistory_PopReturnsLastCommand(t *testing.T) {
	history := command.NewHistory()
	first := &commandStub{name: "first"}
	second := &commandStub{name: "second"}

	history.Push(first)
	history.Push(second)

	got := history.Pop()
	if got != second {
		t.Fatalf("got command %v, want %v", got, second)
	}
	if history.Len() != 1 {
		t.Fatalf("got history length %d, want 1", history.Len())
	}
}
