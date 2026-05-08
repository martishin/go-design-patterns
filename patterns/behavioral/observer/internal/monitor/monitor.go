package monitor

import (
	"errors"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

var (
	ErrEmptyPath   = errors.New("file watcher: path is empty")
	ErrNilObserver = errors.New("file watcher: nil observer")
)

type FileMonitor struct {
	observers []filewatcher.Observer
	revision  int
}

func NewFileMonitor() *FileMonitor {
	return &FileMonitor{}
}

func (m *FileMonitor) Subscribe(observer filewatcher.Observer) error {
	if observer == nil {
		return ErrNilObserver
	}

	for _, existingObserver := range m.observers {
		if existingObserver.ID() == observer.ID() {
			return nil
		}
	}

	m.observers = append(m.observers, observer)
	return nil
}

func (m *FileMonitor) Unsubscribe(observer filewatcher.Observer) {
	if observer == nil {
		return
	}

	for index, existingObserver := range m.observers {
		if existingObserver.ID() == observer.ID() {
			m.observers = append(m.observers[:index], m.observers[index+1:]...)
			return
		}
	}
}

func (m *FileMonitor) Notify(event filewatcher.Event) {
	for _, observer := range m.observers {
		observer.Update(event)
	}
}

func (m *FileMonitor) Write(path string) error {
	return m.publish(path, "write")
}

func (m *FileMonitor) Delete(path string) error {
	return m.publish(path, "delete")
}

func (m *FileMonitor) SubscriberCount() int {
	return len(m.observers)
}

func (m *FileMonitor) publish(path string, operation string) error {
	if path == "" {
		return ErrEmptyPath
	}

	m.revision++
	m.Notify(filewatcher.Event{
		Path:      path,
		Revision:  m.revision,
		Operation: operation,
	})

	return nil
}
