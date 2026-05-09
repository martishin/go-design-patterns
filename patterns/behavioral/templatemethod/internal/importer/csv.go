package importer

import (
	"encoding/csv"
	"errors"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

var ErrInvalidCSV = errors.New("document import: invalid csv data")

type CSVImporter struct {
	source string
	data   string
}

func NewCSVImporter(source string, data string) (*CSVImporter, error) {
	if source == "" {
		return nil, ErrEmptySource
	}

	return &CSVImporter{source: source, data: data}, nil
}

func (i *CSVImporter) Source() string {
	return i.source
}

func (i *CSVImporter) Format() string {
	return "CSV"
}

func (i *CSVImporter) Open() (string, error) {
	if i.data == "" {
		return "", ErrEmptyImportData
	}

	return i.data, nil
}

func (i *CSVImporter) Extract(raw string) ([]document.Record, error) {
	rows, err := csv.NewReader(strings.NewReader(raw)).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, ErrInvalidCSV
	}

	records := make([]document.Record, 0, len(rows)-1)
	for _, row := range rows[1:] {
		if len(row) < 2 {
			return nil, ErrInvalidCSV
		}

		records = append(records, document.Record{ID: row[0], Name: row[1]})
	}

	return records, nil
}
