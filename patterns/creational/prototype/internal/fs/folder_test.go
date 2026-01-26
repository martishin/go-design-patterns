package fs_test

import (
	"bytes"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/prototype/internal/fs"
)

func TestFolder_Print(t *testing.T) {
	// given
	folder := buildSampleFolderTree()

	// when
	var buf bytes.Buffer
	folder.Print(&buf, " ")

	// then
	got := buf.String()
	want := "" +
		"Folder2\n" +
		" Folder1\n" +
		"  File1\n" +
		" File2\n" +
		" File3\n"

	if got != want {
		t.Fatalf("unexpected output:\n--- got ---\n%q\n--- want ---\n%q", got, want)
	}
}

func TestFolder_Clone(t *testing.T) {
	// given
	folder := buildSampleFolderTree()

	// when
	clonedFolder := folder.Clone()

	// then
	var buf bytes.Buffer
	clonedFolder.Print(&buf, " ")

	got := buf.String()
	want := "" +
		"Folder2_clone\n" +
		" Folder1_clone\n" +
		"  File1_clone\n" +
		" File2_clone\n" +
		" File3_clone\n"

	if got != want {
		t.Fatalf("unexpected cloned output:\n--- got ---\n%q\n--- want ---\n%q", got, want)
	}
}

func buildSampleFolderTree() *fs.Folder {
	file1 := fs.NewFile("File1")
	file2 := fs.NewFile("File2")
	file3 := fs.NewFile("File3")

	folder1 := fs.NewFolder("Folder1", file1)
	folder2 := fs.NewFolder("Folder2", folder1, file2, file3)

	return folder2
}
