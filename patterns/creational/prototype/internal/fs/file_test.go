package fs_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/prototype/internal/fs"
)

func TestFile_Clone(t *testing.T) {
	// given
	file := fs.NewFile("File1")

	// when
	clonedFile := file.Clone()
	var buf bytes.Buffer
	clonedFile.Print(&buf, "")

	// then
	got := buf.String()
	want := "File1_clone\n"

	if got != want {
		t.Fatalf("unexpected cloned file name: got %q, want %q", got, want)
	}
}
