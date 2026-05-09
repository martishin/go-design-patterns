package importer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

var (
	ErrNilImporter     = errors.New("document import: nil importer")
	ErrEmptySource     = errors.New("document import: source is empty")
	ErrInvalidRecord   = errors.New("document import: record must have id and name")
	ErrEmptyImportData = errors.New("document import: data is empty")
)

type Pipeline struct {
	importer document.Importer
	stored   []document.Record
}

func NewPipeline(importer document.Importer) (*Pipeline, error) {
	if importer == nil {
		return nil, ErrNilImporter
	}

	return &Pipeline{importer: importer}, nil
}

func (p *Pipeline) Import() (document.Report, error) {
	var steps []string

	raw, err := p.importer.Open()
	if err != nil {
		return document.Report{}, err
	}
	steps = append(steps, fmt.Sprintf("open %s", p.importer.Source()))

	records, err := p.importer.Extract(raw)
	if err != nil {
		return document.Report{}, err
	}
	steps = append(steps, fmt.Sprintf("extract %s records", p.importer.Format()))

	records = normalize(records)
	steps = append(steps, fmt.Sprintf("normalize %d records", len(records)))

	if err := validate(records); err != nil {
		return document.Report{}, err
	}
	steps = append(steps, fmt.Sprintf("validate %d records", len(records)))

	p.store(records)
	steps = append(steps, fmt.Sprintf("store %d records", len(records)))

	return document.Report{
		Source:  p.importer.Source(),
		Format:  p.importer.Format(),
		Steps:   append([]string(nil), steps...),
		Records: append([]document.Record(nil), records...),
	}, nil
}

func (p *Pipeline) StoredRecords() []document.Record {
	return append([]document.Record(nil), p.stored...)
}

func (p *Pipeline) store(records []document.Record) {
	p.stored = append([]document.Record(nil), records...)
}

func normalize(records []document.Record) []document.Record {
	normalized := make([]document.Record, 0, len(records))

	for _, record := range records {
		normalized = append(normalized, document.Record{
			ID:   strings.TrimSpace(record.ID),
			Name: strings.TrimSpace(record.Name),
		})
	}

	return normalized
}

func validate(records []document.Record) error {
	for _, record := range records {
		if record.ID == "" || record.Name == "" {
			return ErrInvalidRecord
		}
	}

	return nil
}
