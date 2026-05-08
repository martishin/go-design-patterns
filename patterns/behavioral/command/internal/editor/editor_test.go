package editor_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
)

func TestEditor_InsertAddsTextAtPosition(t *testing.T) {
	document := editor.NewEditor("deploy api")

	err := document.Insert(6, " payments")
	if err != nil {
		t.Fatalf("Insert() returned error: %v", err)
	}

	got := document.Text()
	want := "deploy payments api"
	if got != want {
		t.Fatalf("got text %q, want %q", got, want)
	}
}

func TestEditor_InsertReturnsErrInvalidRange(t *testing.T) {
	document := editor.NewEditor("deploy api")

	err := document.Insert(99, " payments")
	if !errors.Is(err, editor.ErrInvalidRange) {
		t.Fatalf("got error %v, want %v", err, editor.ErrInvalidRange)
	}
}

func TestEditor_DeleteRemovesTextAndReturnsDeletedText(t *testing.T) {
	document := editor.NewEditor("deploy payments api")

	deleted, err := document.Delete(7, 15)
	if err != nil {
		t.Fatalf("Delete() returned error: %v", err)
	}

	if deleted != "payments" {
		t.Fatalf("got deleted text %q, want %q", deleted, "payments")
	}
	if document.Text() != "deploy  api" {
		t.Fatalf("got text %q, want %q", document.Text(), "deploy  api")
	}
}

func TestEditor_ReplaceChangesTextInRange(t *testing.T) {
	document := editor.NewEditor("deploy api")

	err := document.Replace(0, 6, "release")
	if err != nil {
		t.Fatalf("Replace() returned error: %v", err)
	}

	got := document.Text()
	want := "release api"
	if got != want {
		t.Fatalf("got text %q, want %q", got, want)
	}
}

func TestEditor_SelectionReturnsSelectedText(t *testing.T) {
	document := editor.NewEditor("deploy api")

	got, err := document.Selection(0, 6)
	if err != nil {
		t.Fatalf("Selection() returned error: %v", err)
	}

	if got != "deploy" {
		t.Fatalf("got selection %q, want %q", got, "deploy")
	}
}

func TestEditor_SetClipboardUpdatesClipboard(t *testing.T) {
	document := editor.NewEditor("deploy api")

	document.SetClipboard("deploy")

	if document.Clipboard() != "deploy" {
		t.Fatalf("got clipboard %q, want %q", document.Clipboard(), "deploy")
	}
}
