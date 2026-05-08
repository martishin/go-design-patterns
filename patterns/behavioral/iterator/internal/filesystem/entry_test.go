package filesystem_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/iterator/internal/filesystem"
	iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"
)

func TestNewDirectory_ReturnsErrorForEmptyName(t *testing.T) {
	_, err := filesystem.NewDirectory("")

	if !errors.Is(err, filesystem.ErrEmptyEntryName) {
		t.Fatalf("got error %v, want %v", err, filesystem.ErrEmptyEntryName)
	}
}

func TestNewFile_ReturnsFileEntry(t *testing.T) {
	file, err := filesystem.NewFile("main.go")
	if err != nil {
		t.Fatalf("NewFile() returned error: %v", err)
	}

	var entry iterator.Entry = file
	if entry.Name() != "main.go" {
		t.Fatalf("got name %q, want %q", entry.Name(), "main.go")
	}
	if entry.Path() != "/main.go" {
		t.Fatalf("got path %q, want %q", entry.Path(), "/main.go")
	}
	if entry.IsDirectory() {
		t.Fatal("got directory=true, want false")
	}
}

func TestEntry_AddUpdatesChildPaths(t *testing.T) {
	mainGo, err := filesystem.NewFile("main.go")
	if err != nil {
		t.Fatalf("NewFile() returned error: %v", err)
	}

	cmd, err := filesystem.NewDirectory("cmd", mainGo)
	if err != nil {
		t.Fatalf("NewDirectory() returned error: %v", err)
	}

	app, err := filesystem.NewDirectory("app")
	if err != nil {
		t.Fatalf("NewDirectory() returned error: %v", err)
	}

	if err := app.Add(cmd); err != nil {
		t.Fatalf("Add() returned error: %v", err)
	}

	if cmd.Path() != "/app/cmd" {
		t.Fatalf("got cmd path %q, want %q", cmd.Path(), "/app/cmd")
	}
	if mainGo.Path() != "/app/cmd/main.go" {
		t.Fatalf("got file path %q, want %q", mainGo.Path(), "/app/cmd/main.go")
	}
}

func TestEntry_AddReturnsErrorForNilChild(t *testing.T) {
	directory, err := filesystem.NewDirectory("app")
	if err != nil {
		t.Fatalf("NewDirectory() returned error: %v", err)
	}

	err = directory.Add(nil)

	if !errors.Is(err, filesystem.ErrNilEntry) {
		t.Fatalf("got error %v, want %v", err, filesystem.ErrNilEntry)
	}
}

func TestEntry_AddReturnsErrorForFile(t *testing.T) {
	file, err := filesystem.NewFile("main.go")
	if err != nil {
		t.Fatalf("NewFile() returned error: %v", err)
	}

	child, err := filesystem.NewFile("child.go")
	if err != nil {
		t.Fatalf("NewFile() returned error: %v", err)
	}

	err = file.Add(child)

	if !errors.Is(err, filesystem.ErrNotDirectory) {
		t.Fatalf("got error %v, want %v", err, filesystem.ErrNotDirectory)
	}
}
