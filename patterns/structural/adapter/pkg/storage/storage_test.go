package storage_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/adapter/pkg/storage"
)

type uploaderSpy struct {
	gotPath string
	gotData []byte
	err     error
}

func (s *uploaderSpy) Upload(path string, data []byte) error {
	s.gotPath = path
	s.gotData = append([]byte(nil), data...)
	return s.err
}

func TestBackupService_BackupDelegatesUploader(t *testing.T) {
	spy := &uploaderSpy{}
	service := storage.NewBackupService(spy)

	path := "/snapshots/1"
	data := []byte("snapshot1")

	err := service.Backup(path, data)
	if err != nil {
		t.Fatalf("Backup() returned error: %v", err)
	}
	if spy.gotPath != path {
		t.Fatalf("got path %q, want %q", spy.gotPath, path)
	}
	if !bytes.Equal(spy.gotData, data) {
		t.Fatalf("got data %q, want %q", spy.gotData, data)
	}
}

func TestBackupService_BackupReturnsUploaderError(t *testing.T) {
	wantErr := errors.New("upload failed")
	spy := &uploaderSpy{err: wantErr}
	service := storage.NewBackupService(spy)

	err := service.Backup("/snapshots/1", []byte("snapshot1"))
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}
