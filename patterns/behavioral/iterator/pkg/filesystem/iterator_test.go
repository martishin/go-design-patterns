package filesystem_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"
)

type entryStub struct{}

func (s entryStub) Name() string {
	return "entry"
}

func (s entryStub) Path() string {
	return "/entry"
}

func (s entryStub) IsDirectory() bool {
	return false
}

type iteratorStub struct{}

func (s iteratorStub) HasNext() bool {
	return false
}

func (s iteratorStub) Next() filesystem.Entry {
	return nil
}

type collectionStub struct{}

func (s collectionStub) CreateDepthFirstIterator() filesystem.Iterator {
	return iteratorStub{}
}

func (s collectionStub) CreateBreadthFirstIterator() filesystem.Iterator {
	return iteratorStub{}
}

func TestEntry_InterfaceCanBeImplemented(t *testing.T) {
	var entry filesystem.Entry = entryStub{}

	if entry.Name() != "entry" {
		t.Fatalf("got name %q, want %q", entry.Name(), "entry")
	}
	if entry.Path() != "/entry" {
		t.Fatalf("got path %q, want %q", entry.Path(), "/entry")
	}
	if entry.IsDirectory() {
		t.Fatal("got directory=true, want false")
	}
}

func TestIterator_InterfaceCanBeImplemented(t *testing.T) {
	var iterator filesystem.Iterator = iteratorStub{}

	if iterator.HasNext() {
		t.Fatal("got HasNext()=true, want false")
	}
	if iterator.Next() != nil {
		t.Fatal("got next entry, want nil")
	}
}

func TestCollection_InterfaceCanBeImplemented(t *testing.T) {
	var collection filesystem.Collection = collectionStub{}

	if collection.CreateDepthFirstIterator() == nil {
		t.Fatal("got nil depth-first iterator")
	}
	if collection.CreateBreadthFirstIterator() == nil {
		t.Fatal("got nil breadth-first iterator")
	}
}
