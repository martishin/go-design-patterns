package filesystem_test

import (
	"reflect"
	"testing"

	iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"
)

func TestDepthFirstIterator_TraversesTreePreOrder(t *testing.T) {
	tree := newTestFileTree(t)
	iterator := tree.CreateDepthFirstIterator()

	got := collectPaths(iterator)
	want := []string{
		"/app",
		"/app/cmd",
		"/app/cmd/main.go",
		"/app/internal",
		"/app/internal/server.go",
		"/app/internal/handler.go",
		"/app/README.md",
	}

	assertPaths(t, got, want)
}

func TestDepthFirstIterator_ReturnsNilWhenExhausted(t *testing.T) {
	tree := newTestFileTree(t)
	iterator := tree.CreateDepthFirstIterator()

	collectPaths(iterator)

	if iterator.Next() != nil {
		t.Fatal("got next entry, want nil")
	}
}

func collectPaths(iterator iterator.Iterator) []string {
	var paths []string

	for iterator.HasNext() {
		paths = append(paths, iterator.Next().Path())
	}

	return paths
}

func assertPaths(t *testing.T, got []string, want []string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got paths %v, want %v", got, want)
	}
}
