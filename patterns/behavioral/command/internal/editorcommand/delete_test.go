package editorcommand_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editorcommand"
)

func TestDeleteTextCommand_ExecuteDeletesText(t *testing.T) {
	document := editor.NewEditor("deploy payments api")
	cmd := editorcommand.NewDeleteTextCommand(document, 7, 15)

	changed, err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	if !changed {
		t.Fatal("got changed=false, want true")
	}
	if document.Text() != "deploy  api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy  api")
	}
}

func TestDeleteTextCommand_DescriptionIncludesDeletedRange(t *testing.T) {
	document := editor.NewEditor("deploy payments api")
	cmd := editorcommand.NewDeleteTextCommand(document, 7, 15)

	want := "delete range [7:15]"
	if cmd.Description() != want {
		t.Fatalf("got description %q, want %q", cmd.Description(), want)
	}
}

func TestDeleteTextCommand_UndoRestoresPreviousText(t *testing.T) {
	document := editor.NewEditor("deploy payments api")
	cmd := editorcommand.NewDeleteTextCommand(document, 7, 15)

	if _, err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	cmd.Undo()

	if document.Text() != "deploy payments api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy payments api")
	}
}
