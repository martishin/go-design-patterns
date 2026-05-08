package editorcommand_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editorcommand"
)

func TestInsertTextCommand_ExecuteInsertsText(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewInsertTextCommand(document, len(document.Text()), " to production")

	changed, err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	if !changed {
		t.Fatal("got changed=false, want true")
	}
	if document.Text() != "deploy api to production" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy api to production")
	}
}

func TestInsertTextCommand_DescriptionIncludesInsertedTextAndPosition(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewInsertTextCommand(document, len(document.Text()), " to production")

	want := "insert \" to production\" at position 10"
	if cmd.Description() != want {
		t.Fatalf("got description %q, want %q", cmd.Description(), want)
	}
}

func TestInsertTextCommand_UndoRestoresPreviousText(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewInsertTextCommand(document, len(document.Text()), " to production")

	if _, err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	cmd.Undo()

	if document.Text() != "deploy api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy api")
	}
}
