package importer_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/internal/importer"
	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

func TestNewCSVImporter_ReturnsErrorForEmptySource(t *testing.T) {
	_, err := importer.NewCSVImporter("", "id,name\n1,Alice")

	if !errors.Is(err, importer.ErrEmptySource) {
		t.Fatalf("got error %v, want %v", err, importer.ErrEmptySource)
	}
}

func TestCSVImporter_ImplementsImporter(t *testing.T) {
	csvImporter := newCSVImporter(t, "id,name\n1,Alice")
	var formatImporter document.Importer = csvImporter

	if formatImporter.Source() != "customers.csv" {
		t.Fatalf("got source %q, want %q", formatImporter.Source(), "customers.csv")
	}
	if formatImporter.Format() != "CSV" {
		t.Fatalf("got format %q, want %q", formatImporter.Format(), "CSV")
	}
}

func TestCSVImporter_OpenReturnsErrorForEmptyData(t *testing.T) {
	csvImporter := newCSVImporter(t, "")

	_, err := csvImporter.Open()

	if !errors.Is(err, importer.ErrEmptyImportData) {
		t.Fatalf("got error %v, want %v", err, importer.ErrEmptyImportData)
	}
}

func TestCSVImporter_ExtractReturnsRecords(t *testing.T) {
	csvImporter := newCSVImporter(t, "id,name\n1,Alice\n2,Bob")

	records, err := csvImporter.Extract("id,name\n1,Alice\n2,Bob")
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}

	want := []document.Record{{ID: "1", Name: "Alice"}, {ID: "2", Name: "Bob"}}
	if !reflect.DeepEqual(records, want) {
		t.Fatalf("got records %v, want %v", records, want)
	}
}

func TestCSVImporter_ExtractReturnsErrorForInvalidCSV(t *testing.T) {
	csvImporter := newCSVImporter(t, "id,name")

	_, err := csvImporter.Extract("id,name")

	if !errors.Is(err, importer.ErrInvalidCSV) {
		t.Fatalf("got error %v, want %v", err, importer.ErrInvalidCSV)
	}
}

func newCSVImporter(t *testing.T, data string) *importer.CSVImporter {
	t.Helper()

	csvImporter, err := importer.NewCSVImporter("customers.csv", data)
	if err != nil {
		t.Fatalf("NewCSVImporter() returned error: %v", err)
	}

	return csvImporter
}
