package importer_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/internal/importer"
	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

type importerStub struct {
	source  string
	format  string
	raw     string
	records []document.Record
	err     error
}

func (i *importerStub) Source() string {
	return i.source
}

func (i *importerStub) Format() string {
	return i.format
}

func (i *importerStub) Open() (string, error) {
	return i.raw, i.err
}

func (i *importerStub) Extract(_ string) ([]document.Record, error) {
	return i.records, i.err
}

func TestNewPipeline_ReturnsErrorForNilImporter(t *testing.T) {
	_, err := importer.NewPipeline(nil)

	if !errors.Is(err, importer.ErrNilImporter) {
		t.Fatalf("got error %v, want %v", err, importer.ErrNilImporter)
	}
}

func TestPipeline_ImportRunsTemplateSteps(t *testing.T) {
	formatImporter := &importerStub{
		source: "customers.csv",
		format: "CSV",
		raw:    "id,name\n1,Alice",
		records: []document.Record{
			{ID: " 1 ", Name: " Alice "},
			{ID: "2", Name: "Bob"},
		},
	}
	pipeline, err := importer.NewPipeline(formatImporter)
	if err != nil {
		t.Fatalf("NewPipeline() returned error: %v", err)
	}

	report, err := pipeline.Import()
	if err != nil {
		t.Fatalf("Import() returned error: %v", err)
	}

	wantSteps := []string{
		"open customers.csv",
		"extract CSV records",
		"normalize 2 records",
		"validate 2 records",
		"store 2 records",
	}
	if !reflect.DeepEqual(report.Steps, wantSteps) {
		t.Fatalf("got steps %v, want %v", report.Steps, wantSteps)
	}

	wantRecords := []document.Record{{ID: "1", Name: "Alice"}, {ID: "2", Name: "Bob"}}
	if !reflect.DeepEqual(report.Records, wantRecords) {
		t.Fatalf("got records %v, want %v", report.Records, wantRecords)
	}
	if !reflect.DeepEqual(pipeline.StoredRecords(), wantRecords) {
		t.Fatalf("got stored records %v, want %v", pipeline.StoredRecords(), wantRecords)
	}
}

func TestPipeline_ImportReturnsValidationError(t *testing.T) {
	formatImporter := &importerStub{
		source:  "customers.csv",
		format:  "CSV",
		raw:     "id,name\n1,",
		records: []document.Record{{ID: "1"}},
	}
	pipeline, err := importer.NewPipeline(formatImporter)
	if err != nil {
		t.Fatalf("NewPipeline() returned error: %v", err)
	}

	_, err = pipeline.Import()

	if !errors.Is(err, importer.ErrInvalidRecord) {
		t.Fatalf("got error %v, want %v", err, importer.ErrInvalidRecord)
	}
}

func TestPipeline_StoredRecordsReturnsCopy(t *testing.T) {
	formatImporter := &importerStub{
		source:  "customers.csv",
		format:  "CSV",
		raw:     "id,name\n1,Alice",
		records: []document.Record{{ID: "1", Name: "Alice"}},
	}
	pipeline, err := importer.NewPipeline(formatImporter)
	if err != nil {
		t.Fatalf("NewPipeline() returned error: %v", err)
	}
	if _, err := pipeline.Import(); err != nil {
		t.Fatalf("Import() returned error: %v", err)
	}

	records := pipeline.StoredRecords()
	records[0] = document.Record{ID: "changed", Name: "changed"}

	want := []document.Record{{ID: "1", Name: "Alice"}}
	if !reflect.DeepEqual(pipeline.StoredRecords(), want) {
		t.Fatalf("stored records were mutated through returned slice: %v", pipeline.StoredRecords())
	}
}
