package filesystem

import iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"

type DepthFirstIterator struct {
	stack []*Entry
}

func NewDepthFirstIterator(root *Entry) *DepthFirstIterator {
	if root == nil {
		return &DepthFirstIterator{}
	}

	return &DepthFirstIterator{stack: []*Entry{root}}
}

func (i *DepthFirstIterator) HasNext() bool {
	return len(i.stack) > 0
}

func (i *DepthFirstIterator) Next() iterator.Entry {
	if !i.HasNext() {
		return nil
	}

	lastIndex := len(i.stack) - 1
	entry := i.stack[lastIndex]
	i.stack = i.stack[:lastIndex]

	for childIndex := len(entry.children) - 1; childIndex >= 0; childIndex-- {
		i.stack = append(i.stack, entry.children[childIndex])
	}

	return entry
}
