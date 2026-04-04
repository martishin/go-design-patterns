package main

import (
	"github.com/martishin/go-design-patterns/patterns/structural/adapter/internal/bytestore"
	"github.com/martishin/go-design-patterns/patterns/structural/adapter/internal/streamadapter"
	"github.com/martishin/go-design-patterns/patterns/structural/adapter/internal/streamstore"
	"github.com/martishin/go-design-patterns/patterns/structural/adapter/pkg/storage"
)

func main() {
	byteStore := bytestore.NewByteStore()
	streamStore := streamstore.NewStreamStore()
	streamStoreAdapter := streamadapter.NewStreamStoreAdapter(streamStore)

	byteStoreUploader := storage.NewBackupService(byteStore)
	streamStoreUploader := storage.NewBackupService(streamStoreAdapter)

	_ = byteStoreUploader.Backup("/snapshots/1", []byte("snapshot1"))
	_ = streamStoreUploader.Backup("/snapshots/2", []byte("snapshot2"))
}
