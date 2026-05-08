package filesystem

import iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"

type BreadthFirstIterator struct {
	queue []*Entry
}

func NewBreadthFirstIterator(root *Entry) *BreadthFirstIterator {
	if root == nil {
		return &BreadthFirstIterator{}
	}

	return &BreadthFirstIterator{queue: []*Entry{root}}
}

func (i *BreadthFirstIterator) HasNext() bool {
	return len(i.queue) > 0
}

func (i *BreadthFirstIterator) Next() iterator.Entry {
	if !i.HasNext() {
		return nil
	}

	entry := i.queue[0]
	i.queue = i.queue[1:]
	i.queue = append(i.queue, entry.children...)

	return entry
}
