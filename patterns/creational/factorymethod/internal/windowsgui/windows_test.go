package windowsgui_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/internal/windowsgui"
)

func TestWindowsDialog(t *testing.T) {
	dialog := windowsgui.New()

	var buf bytes.Buffer
	dialog.Render(&buf)
	dialog.Refresh(&buf)

	got := buf.String()
	want := "" +
		"Drawing a Windows button\n" +
		"Click! Hello, Windows!\n" +
		"Dialog - Refresh\n"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
