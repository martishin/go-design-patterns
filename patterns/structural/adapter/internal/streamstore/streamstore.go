package streamstore

import (
	"fmt"
	"io"
)

type StreamStore struct {
}

func (s *StreamStore) Put(path string, body io.Reader) error {
	content, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	fmt.Printf("uploading stream data \"%s\" to \"%s\"\n", content, path)
	return nil
}

func NewStreamStore() *StreamStore {
	return &StreamStore{}
}
