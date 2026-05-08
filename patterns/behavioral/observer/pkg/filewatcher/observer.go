package filewatcher

type Event struct {
	Path      string
	Revision  int
	Operation string
}

type Observer interface {
	ID() string
	Update(event Event)
}

type Subject interface {
	Subscribe(observer Observer) error
	Unsubscribe(observer Observer)
	Notify(event Event)
}
