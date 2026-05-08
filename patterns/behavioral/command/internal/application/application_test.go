package application_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/application"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type commandSpy struct {
	changed bool
	err     error
	undone  bool
}

func (s *commandSpy) Name() string {
	return "spy command"
}

func (s *commandSpy) Description() string {
	return "spy command description"
}

func (s *commandSpy) Execute() (bool, error) {
	return s.changed, s.err
}

func (s *commandSpy) Undo() {
	s.undone = true
}

func TestApplication_ExecutePushesMutatingCommandToHistory(t *testing.T) {
	app := application.NewApplication(command.NewHistory())
	cmd := &commandSpy{changed: true}

	err := app.Execute(cmd)
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	if app.HistoryLen() != 1 {
		t.Fatalf("got history length %d, want 1", app.HistoryLen())
	}
}

func TestApplication_ExecuteSkipsNonMutatingCommand(t *testing.T) {
	app := application.NewApplication(command.NewHistory())
	cmd := &commandSpy{changed: false}

	err := app.Execute(cmd)
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	if app.HistoryLen() != 0 {
		t.Fatalf("got history length %d, want 0", app.HistoryLen())
	}
}

func TestApplication_ExecuteReturnsCommandError(t *testing.T) {
	app := application.NewApplication(command.NewHistory())
	wantErr := errors.New("command failed")
	cmd := &commandSpy{err: wantErr}

	err := app.Execute(cmd)
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
	if app.HistoryLen() != 0 {
		t.Fatalf("got history length %d, want 0", app.HistoryLen())
	}
}

func TestApplication_UndoRunsLastCommandUndo(t *testing.T) {
	history := command.NewHistory()
	app := application.NewApplication(history)
	cmd := &commandSpy{changed: true}

	if err := app.Execute(cmd); err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	undoneCommand := app.Undo()

	if !cmd.undone {
		t.Fatal("expected Undo() to run on last command")
	}
	if undoneCommand != cmd {
		t.Fatalf("got undone command %v, want %v", undoneCommand, cmd)
	}
	if app.HistoryLen() != 0 {
		t.Fatalf("got history length %d, want 0", app.HistoryLen())
	}
}
