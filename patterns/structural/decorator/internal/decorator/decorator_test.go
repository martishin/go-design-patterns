package decorator_test

import (
	"bytes"
	"errors"
	"path/filepath"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/decorator"
	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/filedata"
	"github.com/martishin/go-design-patterns/patterns/structural/decorator/pkg/datasource"
)

func TestDataSourceDecorator_WriteDelegatesToWrappee(t *testing.T) {
	spy := &dataSourceSpy{}
	wrapper, err := decorator.NewDataSourceDecorator(spy)
	if err != nil {
		t.Fatalf("NewDataSourceDecorator() returned error: %v", err)
	}

	payload := []byte("salary-records")
	if err := wrapper.Write(payload); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if spy.writeCalls != 1 {
		t.Fatalf("got %d write calls, want 1", spy.writeCalls)
	}
	if !bytes.Equal(spy.gotWriteData, payload) {
		t.Fatalf("got write payload %q, want %q", spy.gotWriteData, payload)
	}
}

func TestDataSourceDecorator_ReadDelegatesToWrappee(t *testing.T) {
	want := []byte("salary-records")
	spy := &dataSourceSpy{readData: want}
	wrapper, err := decorator.NewDataSourceDecorator(spy)
	if err != nil {
		t.Fatalf("NewDataSourceDecorator() returned error: %v", err)
	}

	got, err := wrapper.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if spy.readCalls != 1 {
		t.Fatalf("got %d read calls, want 1", spy.readCalls)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got read payload %q, want %q", got, want)
	}
}

func TestDataSourceDecorator_NewReturnsErrNoDataSourceForNilWrappee(t *testing.T) {
	_, err := decorator.NewDataSourceDecorator(nil)
	if !errors.Is(err, decorator.ErrNoDataSource) {
		t.Fatalf("got error %v, want %v", err, decorator.ErrNoDataSource)
	}
}

func TestDataSourceDecorator_WriteReturnsErrNoDataSource(t *testing.T) {
	var wrapper decorator.DataSourceDecorator

	err := wrapper.Write([]byte("salary-records"))
	if !errors.Is(err, decorator.ErrNoDataSource) {
		t.Fatalf("got error %v, want %v", err, decorator.ErrNoDataSource)
	}
}

func TestDataSourceDecorator_ReadReturnsErrNoDataSource(t *testing.T) {
	var wrapper decorator.DataSourceDecorator

	_, err := wrapper.Read()
	if !errors.Is(err, decorator.ErrNoDataSource) {
		t.Fatalf("got error %v, want %v", err, decorator.ErrNoDataSource)
	}
}

func TestDataSourceDecorator_NewReturnsErrNoDataSourceForTypedNilWrappee(t *testing.T) {
	var typedNil *dataSourceSpy
	var source datasource.DataSource = typedNil

	_, err := decorator.NewDataSourceDecorator(source)
	if !errors.Is(err, decorator.ErrNoDataSource) {
		t.Fatalf("got error %v, want %v", err, decorator.ErrNoDataSource)
	}
}

func TestDecoratorStack_FileRoundTripRestoresOriginalBytes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "salary.dat")
	key := []byte("0123456789abcdef0123456789abcdef")
	want := []byte("name,salary\nAlice,120000\nBob,95000")

	fileSource := filedata.NewFileDataSource(path)

	encryptedSource, err := decorator.NewEncryptionDecorator(fileSource, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	compressedAndEncryptedSource, err := decorator.NewCompressionDecorator(encryptedSource)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	if err := compressedAndEncryptedSource.Write(want); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	got, err := compressedAndEncryptedSource.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got payload %q, want %q", got, want)
	}
}

func TestDecoratorStack_DecoratedWriteDoesNotStorePlaintext(t *testing.T) {
	path := filepath.Join(t.TempDir(), "salary.dat")
	key := []byte("0123456789abcdef0123456789abcdef")
	payload := []byte("name,salary\nAlice,120000\nBob,95000")

	fileSource := filedata.NewFileDataSource(path)

	encryptedSource, err := decorator.NewEncryptionDecorator(fileSource, key)
	if err != nil {
		t.Fatalf("NewEncryptionDecorator() returned error: %v", err)
	}

	compressedAndEncryptedSource, err := decorator.NewCompressionDecorator(encryptedSource)
	if err != nil {
		t.Fatalf("NewCompressionDecorator() returned error: %v", err)
	}

	if err := compressedAndEncryptedSource.Write(payload); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	stored, err := fileSource.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if bytes.Equal(stored, payload) {
		t.Fatal("stored bytes unexpectedly match plaintext payload")
	}
	if bytes.Contains(stored, []byte("Alice")) {
		t.Fatal("stored bytes unexpectedly contain plaintext data")
	}
}

type dataSourceSpy struct {
	writeCalls int
	readCalls  int

	gotWriteData []byte
	readData     []byte

	writeErr error
	readErr  error
}

func (s *dataSourceSpy) Write(data []byte) error {
	s.writeCalls++
	s.gotWriteData = append([]byte(nil), data...)
	return s.writeErr
}

func (s *dataSourceSpy) Read() ([]byte, error) {
	s.readCalls++
	return append([]byte(nil), s.readData...), s.readErr
}
