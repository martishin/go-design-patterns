package editorcommand_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editorcommand"
)

func TestReplaceTextCommand_ExecuteReplacesText(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewReplaceTextCommand(document, 0, 6, "release")

	changed, err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	if !changed {
		t.Fatal("got changed=false, want true")
	}
	if document.Text() != "release api" {
		t.Fatalf("got text %q, want %q", document.Text(), "release api")
	}
}

func TestReplaceTextCommand_DescriptionIncludesRangeAndReplacement(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewReplaceTextCommand(document, 0, 6, "release")

	want := "replace range [0:6] with \"release\""
	if cmd.Description() != want {
		t.Fatalf("got description %q, want %q", cmd.Description(), want)
	}
}

func TestReplaceTextCommand_UndoRestoresPreviousText(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewReplaceTextCommand(document, 0, 6, "release")

	if _, err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	cmd.Undo()

	if document.Text() != "deploy api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy api")
	}
}
