package editorcommand

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
)

func TestBaseCommand_UndoRestoresBackup(t *testing.T) {
	document := editor.NewEditor("deploy api")
	base := newBaseCommand(document)

	base.saveBackup()
	document.SetText("release api")
	base.Undo()

	if document.Text() != "deploy api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy api")
	}
}
