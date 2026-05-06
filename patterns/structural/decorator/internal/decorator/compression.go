package decorator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/pkg/datasource"
)

type CompressionDecorator struct {
	*DataSourceDecorator
}

func NewCompressionDecorator(wrappee datasource.DataSource) (datasource.DataSource, error) {
	base, err := NewDataSourceDecorator(wrappee)
	if err != nil {
		return nil, err
	}

	return &CompressionDecorator{DataSourceDecorator: base}, nil
}

func (d *CompressionDecorator) Write(data []byte) error {
	if d == nil || d.DataSourceDecorator == nil {
		return ErrNoDataSource
	}

	compressed, err := compressData(data)
	if err != nil {
		return fmt.Errorf("compress data: %w", err)
	}

	return d.DataSourceDecorator.Write(compressed)
}

func (d *CompressionDecorator) Read() ([]byte, error) {
	if d == nil || d.DataSourceDecorator == nil {
		return nil, ErrNoDataSource
	}

	compressed, err := d.DataSourceDecorator.Read()
	if err != nil {
		return nil, err
	}

	data, err := decompressData(compressed)
	if err != nil {
		return nil, fmt.Errorf("decompress data: %w", err)
	}

	return data, nil
}

func compressData(data []byte) ([]byte, error) {
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

func decompressData(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	result, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return result, nil
}
