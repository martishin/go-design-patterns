package filesystem

import "errors"

var (
	ErrEmptyEntryName = errors.New("filesystem: entry name is empty")
	ErrNilEntry       = errors.New("filesystem: nil entry")
	ErrNotDirectory   = errors.New("filesystem: entry is not a directory")
)

type Entry struct {
	name        string
	path        string
	isDirectory bool
	children    []*Entry
}

func NewDirectory(name string, children ...*Entry) (*Entry, error) {
	entry, err := newEntry(name, true)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		if err := entry.Add(child); err != nil {
			return nil, err
		}
	}

	return entry, nil
}

func NewFile(name string) (*Entry, error) {
	return newEntry(name, false)
}

func newEntry(name string, isDirectory bool) (*Entry, error) {
	if name == "" {
		return nil, ErrEmptyEntryName
	}

	entry := &Entry{
		name:        name,
		path:        "/" + name,
		isDirectory: isDirectory,
	}

	return entry, nil
}

func (e *Entry) Name() string {
	return e.name
}

func (e *Entry) Path() string {
	return e.path
}

func (e *Entry) IsDirectory() bool {
	return e.isDirectory
}

func (e *Entry) Add(child *Entry) error {
	if child == nil {
		return ErrNilEntry
	}
	if !e.isDirectory {
		return ErrNotDirectory
	}

	child.updatePath(e.path)
	e.children = append(e.children, child)

	return nil
}

func (e *Entry) updatePath(parentPath string) {
	e.path = parentPath + "/" + e.name

	for _, child := range e.children {
		child.updatePath(e.path)
	}
}
