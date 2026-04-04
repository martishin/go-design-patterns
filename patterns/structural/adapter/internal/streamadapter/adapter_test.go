package streamadapter_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/adapter/internal/streamadapter"
)

type putSpy struct {
	gotPath string
	gotBody []byte
	err     error
}

func (s *putSpy) Put(path string, body io.Reader) error {
	s.gotPath = path

	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	s.gotBody = data

	return s.err
}

func TestAdapter_UploadPassesPathAndContentToStore(t *testing.T) {
	spy := &putSpy{}
	uploader := streamadapter.NewStreamStoreAdapter(spy)

	path := "/snapshots/2"
	data := []byte("snapshot2")

	err := uploader.Upload(path, data)
	if err != nil {
		t.Fatalf("Upload() returned error: %v", err)
	}
	if spy.gotPath != path {
		t.Fatalf("got path %q, want %q", spy.gotPath, path)
	}
	if !bytes.Equal(spy.gotBody, data) {
		t.Fatalf("got body %q, want %q", spy.gotBody, data)
	}
}

func TestAdapter_UploadReturnsStoreError(t *testing.T) {
	wantErr := errors.New("put failed")
	spy := &putSpy{err: wantErr}
	uploader := streamadapter.NewStreamStoreAdapter(spy)

	err := uploader.Upload("/snapshots/2", []byte("snapshot2"))
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}
