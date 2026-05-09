package importer

import (
	"encoding/json"
	"errors"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

var ErrInvalidJSON = errors.New("document import: invalid json data")

type JSONImporter struct {
	source string
	data   string
}

func NewJSONImporter(source string, data string) (*JSONImporter, error) {
	if source == "" {
		return nil, ErrEmptySource
	}

	return &JSONImporter{source: source, data: data}, nil
}

func (i *JSONImporter) Source() string {
	return i.source
}

func (i *JSONImporter) Format() string {
	return "JSON"
}

func (i *JSONImporter) Open() (string, error) {
	if i.data == "" {
		return "", ErrEmptyImportData
	}

	return i.data, nil
}

func (i *JSONImporter) Extract(raw string) ([]document.Record, error) {
	var records []document.Record
	if err := json.Unmarshal([]byte(raw), &records); err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, ErrInvalidJSON
	}

	return records, nil
}
