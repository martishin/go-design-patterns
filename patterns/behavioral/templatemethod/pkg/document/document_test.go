package document_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

type importerStub struct{}

func (i importerStub) Source() string {
	return "customers.csv"
}

func (i importerStub) Format() string {
	return "CSV"
}

func (i importerStub) Open() (string, error) {
	return "id,name\n1,Alice", nil
}

func (i importerStub) Extract(_ string) ([]document.Record, error) {
	return []document.Record{{ID: "1", Name: "Alice"}}, nil
}

func TestImporter_InterfaceCanBeImplemented(t *testing.T) {
	var importer document.Importer = importerStub{}

	if importer.Source() != "customers.csv" {
		t.Fatalf("got source %q, want %q", importer.Source(), "customers.csv")
	}
	if importer.Format() != "CSV" {
		t.Fatalf("got format %q, want %q", importer.Format(), "CSV")
	}
}
