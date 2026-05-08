package editorcommand_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editorcommand"
)

func TestCopyTextCommand_ExecuteCopiesSelection(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewCopyTextCommand(document, 0, 6)

	changed, err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
	if changed {
		t.Fatal("got changed=true, want false")
	}
	if document.Clipboard() != "deploy" {
		t.Fatalf("got clipboard %q, want %q", document.Clipboard(), "deploy")
	}
	if document.Text() != "deploy api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy api")
	}
}

func TestCopyTextCommand_DescriptionIncludesSelectionRange(t *testing.T) {
	document := editor.NewEditor("deploy api")
	cmd := editorcommand.NewCopyTextCommand(document, 0, 6)

	want := "copy range [0:6] to clipboard"
	if cmd.Description() != want {
		t.Fatalf("got description %q, want %q", cmd.Description(), want)
	}
}
