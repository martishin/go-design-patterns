package filesystem

import iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"

type FileTree struct {
	root *Entry
}

func NewFileTree(root *Entry) (*FileTree, error) {
	if root == nil {
		return nil, ErrNilEntry
	}

	return &FileTree{root: root}, nil
}

func (t *FileTree) CreateDepthFirstIterator() iterator.Iterator {
	return NewDepthFirstIterator(t.root)
}

func (t *FileTree) CreateBreadthFirstIterator() iterator.Iterator {
	return NewBreadthFirstIterator(t.root)
}
