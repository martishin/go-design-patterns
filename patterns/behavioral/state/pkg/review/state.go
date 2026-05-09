package review

type State interface {
	Name() string
	Open() error
	Approve() error
	RequestChanges() error
	Merge() error
}
