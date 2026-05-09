package importer_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/internal/importer"
	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

func TestNewJSONImporter_ReturnsErrorForEmptySource(t *testing.T) {
	_, err := importer.NewJSONImporter("", `[{"id":"1","name":"Alice"}]`)

	if !errors.Is(err, importer.ErrEmptySource) {
		t.Fatalf("got error %v, want %v", err, importer.ErrEmptySource)
	}
}

func TestJSONImporter_ImplementsImporter(t *testing.T) {
	jsonImporter := newJSONImporter(t, `[{"id":"1","name":"Alice"}]`)
	var formatImporter document.Importer = jsonImporter

	if formatImporter.Source() != "customers.json" {
		t.Fatalf("got source %q, want %q", formatImporter.Source(), "customers.json")
	}
	if formatImporter.Format() != "JSON" {
		t.Fatalf("got format %q, want %q", formatImporter.Format(), "JSON")
	}
}

func TestJSONImporter_OpenReturnsErrorForEmptyData(t *testing.T) {
	jsonImporter := newJSONImporter(t, "")

	_, err := jsonImporter.Open()

	if !errors.Is(err, importer.ErrEmptyImportData) {
		t.Fatalf("got error %v, want %v", err, importer.ErrEmptyImportData)
	}
}

func TestJSONImporter_ExtractReturnsRecords(t *testing.T) {
	jsonImporter := newJSONImporter(t, `[{"id":"1","name":"Alice"},{"id":"2","name":"Bob"}]`)

	records, err := jsonImporter.Extract(`[{"id":"1","name":"Alice"},{"id":"2","name":"Bob"}]`)
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}

	want := []document.Record{{ID: "1", Name: "Alice"}, {ID: "2", Name: "Bob"}}
	if !reflect.DeepEqual(records, want) {
		t.Fatalf("got records %v, want %v", records, want)
	}
}

func TestJSONImporter_ExtractReturnsErrorForInvalidJSON(t *testing.T) {
	jsonImporter := newJSONImporter(t, `[]`)

	_, err := jsonImporter.Extract(`[]`)

	if !errors.Is(err, importer.ErrInvalidJSON) {
		t.Fatalf("got error %v, want %v", err, importer.ErrInvalidJSON)
	}
}

func newJSONImporter(t *testing.T, data string) *importer.JSONImporter {
	t.Helper()

	jsonImporter, err := importer.NewJSONImporter("customers.json", data)
	if err != nil {
		t.Fatalf("NewJSONImporter() returned error: %v", err)
	}

	return jsonImporter
}
