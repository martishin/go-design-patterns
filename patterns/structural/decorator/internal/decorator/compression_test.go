package decorator_test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/decorator"
)

func TestCompressionDecorator_WriteCompressesBeforeDelegating(t *testing.T) {
	spy := &dataSourceSpy{}
	compressedDataSource, err := decorator.NewCompressionDecorator(spy)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	payload := []byte("salary records salary records salary records")
	if err := compressedDataSource.Write(payload); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if spy.writeCalls != 1 {
		t.Fatalf("got %d write calls, want 1", spy.writeCalls)
	}
	if bytes.Equal(spy.gotWriteData, payload) {
		t.Fatalf("expected compressed bytes to differ from original payload")
	}

	got, err := decompressWithGzip(spy.gotWriteData)
	if err != nil {
		t.Fatalf("decompressWithGzip() returned error: %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("got decompressed payload %q, want %q", got, payload)
	}
}

func TestCompressionDecorator_ReadDecompressesWrappedBytes(t *testing.T) {
	want := []byte("salary records salary records salary records")
	compressed, err := compressWithGzip(want)
	if err != nil {
		t.Fatalf("compressWithGzip() returned error: %v", err)
	}

	spy := &dataSourceSpy{readData: compressed}
	compressedDataSource, err := decorator.NewCompressionDecorator(spy)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	got, err := compressedDataSource.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if spy.readCalls != 1 {
		t.Fatalf("got %d read calls, want 1", spy.readCalls)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got payload %q, want %q", got, want)
	}
}

func TestCompressionDecorator_WriteReturnsWrappedError(t *testing.T) {
	wantErr := errors.New("write failed")
	spy := &dataSourceSpy{writeErr: wantErr}
	compressedDataSource, err := decorator.NewCompressionDecorator(spy)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	err = compressedDataSource.Write([]byte("salary records"))
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestCompressionDecorator_ReadReturnsWrappedError(t *testing.T) {
	wantErr := errors.New("read failed")
	spy := &dataSourceSpy{readErr: wantErr}
	compressedDataSource, err := decorator.NewCompressionDecorator(spy)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	_, err = compressedDataSource.Read()
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestCompressionDecorator_ReadReturnsErrorForInvalidCompressedBytes(t *testing.T) {
	spy := &dataSourceSpy{readData: []byte("not-gzip")}
	compressedDataSource, err := decorator.NewCompressionDecorator(spy)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	_, err = compressedDataSource.Read()
	if err == nil {
		t.Fatal("Read() returned nil error, want non-nil")
	}
}

func compressWithGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decompressWithGzip(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}
