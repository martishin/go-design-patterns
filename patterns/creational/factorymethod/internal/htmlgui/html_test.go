package htmlgui_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/internal/htmlgui"
)

func TestHTMLDialog_GoldenOutput(t *testing.T) {
	dialog := htmlgui.New()

	var buf bytes.Buffer
	dialog.Render(&buf)
	dialog.Refresh(&buf)

	got := buf.String()

	want := "<button>Test Button</button>\n" +
		"Click! Button says - 'Hello World!'\n" +
		"Dialog - Refresh\n"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
