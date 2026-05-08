package filesystem

type Entry interface {
	Name() string
	Path() string
	IsDirectory() bool
}

type Iterator interface {
	HasNext() bool
	Next() Entry
}

type Collection interface {
	CreateDepthFirstIterator() Iterator
	CreateBreadthFirstIterator() Iterator
}
