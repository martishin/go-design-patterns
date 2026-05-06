package filedata_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/filedata"
)

func TestFileDataSource_WriteCreatesFileWithExactBytes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "salary.dat")
	source := filedata.NewFileDataSource(path)
	want := []byte("salary-records")

	if err := source.Write(want); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() returned error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got file content %q, want %q", got, want)
	}
}

func TestFileDataSource_ReadReturnsFileContents(t *testing.T) {
	path := filepath.Join(t.TempDir(), "salary.dat")
	want := []byte("salary-records")
	if err := os.WriteFile(path, want, 0o644); err != nil {
		t.Fatalf("WriteFile() returned error: %v", err)
	}

	source := filedata.NewFileDataSource(path)

	got, err := source.Read()
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("got content %q, want %q", got, want)
	}
}

func TestFileDataSource_ReadReturnsErrorForMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.dat")
	source := filedata.NewFileDataSource(path)

	_, err := source.Read()
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("got error %v, want %v", err, os.ErrNotExist)
	}
}
