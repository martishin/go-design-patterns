package bytestore

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/structural/adapter/pkg/storage"
)

type ByteStore struct {
}

func (bs *ByteStore) Upload(path string, data []byte) error {
	fmt.Printf("uploading byte data \"%s\" to \"%s\"\n", data, path)
	return nil
}

func NewByteStore() storage.Uploader {
	return &ByteStore{}
}
