package filesystem_test

import "testing"

func TestBreadthFirstIterator_TraversesTreeByLevel(t *testing.T) {
	tree := newTestFileTree(t)
	iterator := tree.CreateBreadthFirstIterator()

	got := collectPaths(iterator)
	want := []string{
		"/app",
		"/app/cmd",
		"/app/internal",
		"/app/README.md",
		"/app/cmd/main.go",
		"/app/internal/server.go",
		"/app/internal/handler.go",
	}

	assertPaths(t, got, want)
}

func TestBreadthFirstIterator_ReturnsNilWhenExhausted(t *testing.T) {
	tree := newTestFileTree(t)
	iterator := tree.CreateBreadthFirstIterator()

	collectPaths(iterator)

	if iterator.Next() != nil {
		t.Fatal("got next entry, want nil")
	}
}
