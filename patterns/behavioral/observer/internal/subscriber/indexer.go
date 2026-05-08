package subscriber

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

type Indexer struct {
	id      string
	updates []string
}

func NewIndexer(id string) *Indexer {
	return &Indexer{id: id}
}

func (i *Indexer) ID() string {
	return i.id
}

func (i *Indexer) Update(event filewatcher.Event) {
	switch event.Operation {
	case "write":
		i.updates = append(i.updates, fmt.Sprintf("indexed %s at revision %d", event.Path, event.Revision))
	case "delete":
		i.updates = append(i.updates, fmt.Sprintf("removed %s from index at revision %d", event.Path, event.Revision))
	}
}

func (i *Indexer) Updates() []string {
	return append([]string(nil), i.updates...)
}
