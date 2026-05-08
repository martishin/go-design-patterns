package filesystem_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/iterator/internal/filesystem"
	iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"
)

func TestNewFileTree_ReturnsErrorForNilRoot(t *testing.T) {
	_, err := filesystem.NewFileTree(nil)

	if !errors.Is(err, filesystem.ErrNilEntry) {
		t.Fatalf("got error %v, want %v", err, filesystem.ErrNilEntry)
	}
}

func TestFileTree_ImplementsCollection(t *testing.T) {
	tree := newTestFileTree(t)
	var collection iterator.Collection = tree

	if collection.CreateDepthFirstIterator() == nil {
		t.Fatal("got nil depth-first iterator")
	}
	if collection.CreateBreadthFirstIterator() == nil {
		t.Fatal("got nil breadth-first iterator")
	}
}

func TestFileTree_ReturnsFreshIterators(t *testing.T) {
	tree := newTestFileTree(t)

	first := tree.CreateDepthFirstIterator()
	second := tree.CreateDepthFirstIterator()

	if first == second {
		t.Fatal("expected fresh iterator instances")
	}

	firstEntry := first.Next()
	if firstEntry.Path() != "/app" {
		t.Fatalf("got first iterator path %q, want %q", firstEntry.Path(), "/app")
	}

	secondEntry := second.Next()
	if secondEntry.Path() != "/app" {
		t.Fatalf("got second iterator path %q, want %q", secondEntry.Path(), "/app")
	}
}

func newTestFileTree(t *testing.T) *filesystem.FileTree {
	t.Helper()

	mainGo := newTestFile(t, "main.go")
	serverGo := newTestFile(t, "server.go")
	handlerGo := newTestFile(t, "handler.go")
	readme := newTestFile(t, "README.md")
	cmd := newTestDirectory(t, "cmd", mainGo)
	internal := newTestDirectory(t, "internal", serverGo, handlerGo)
	app := newTestDirectory(t, "app", cmd, internal, readme)

	tree, err := filesystem.NewFileTree(app)
	if err != nil {
		t.Fatalf("NewFileTree() returned error: %v", err)
	}

	return tree
}

func newTestFile(t *testing.T, name string) *filesystem.Entry {
	t.Helper()

	file, err := filesystem.NewFile(name)
	if err != nil {
		t.Fatalf("NewFile() returned error: %v", err)
	}

	return file
}

func newTestDirectory(t *testing.T, name string, children ...*filesystem.Entry) *filesystem.Entry {
	t.Helper()

	directory, err := filesystem.NewDirectory(name, children...)
	if err != nil {
		t.Fatalf("NewDirectory() returned error: %v", err)
	}

	return directory
}
