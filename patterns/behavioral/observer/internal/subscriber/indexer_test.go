package subscriber_test

import (
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/internal/subscriber"
	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

func TestIndexer_ImplementsObserver(t *testing.T) {
	var _ filewatcher.Observer = subscriber.NewIndexer("indexer")
}

func TestIndexer_UpdateStoresIndexUpdateForWriteEvent(t *testing.T) {
	indexer := subscriber.NewIndexer("indexer")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}

	indexer.Update(event)

	want := []string{"indexed app/config.yaml at revision 1"}
	if !reflect.DeepEqual(indexer.Updates(), want) {
		t.Fatalf("got updates %v, want %v", indexer.Updates(), want)
	}
}

func TestIndexer_UpdateStoresIndexUpdateForDeleteEvent(t *testing.T) {
	indexer := subscriber.NewIndexer("indexer")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 2, Operation: "delete"}

	indexer.Update(event)

	want := []string{"removed app/config.yaml from index at revision 2"}
	if !reflect.DeepEqual(indexer.Updates(), want) {
		t.Fatalf("got updates %v, want %v", indexer.Updates(), want)
	}
}

func TestIndexer_UpdatesReturnsCopy(t *testing.T) {
	indexer := subscriber.NewIndexer("indexer")
	event := filewatcher.Event{Path: "app/config.yaml", Revision: 1, Operation: "write"}
	indexer.Update(event)

	updates := indexer.Updates()
	updates[0] = "changed"

	if indexer.Updates()[0] != "indexed app/config.yaml at revision 1" {
		t.Fatalf("updates were mutated through returned slice: %v", indexer.Updates())
	}
}
