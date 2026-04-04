package streamadapter

import (
	"bytes"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/adapter/pkg/storage"
)

type putter interface {
	Put(path string, body io.Reader) error
}

type StreamStoreAdapter struct {
	store putter
}

func (s *StreamStoreAdapter) Upload(path string, data []byte) error {
	return s.store.Put(path, bytes.NewReader(data))
}

func NewStreamStoreAdapter(store putter) storage.Uploader {
	return &StreamStoreAdapter{
		store: store,
	}
}
