package decorator

import (
	"errors"
	"reflect"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/pkg/datasource"
)

var ErrNoDataSource = errors.New("decorator: no wrapped data source")

type DataSourceDecorator struct {
	wrapped datasource.DataSource
}

func NewDataSourceDecorator(source datasource.DataSource) (*DataSourceDecorator, error) {
	if isNilDataSource(source) {
		return nil, ErrNoDataSource
	}

	return &DataSourceDecorator{wrapped: source}, nil
}

func (d *DataSourceDecorator) Write(data []byte) error {
	if d == nil || isNilDataSource(d.wrapped) {
		return ErrNoDataSource
	}

	return d.wrapped.Write(data)
}

func (d *DataSourceDecorator) Read() ([]byte, error) {
	if d == nil || isNilDataSource(d.wrapped) {
		return nil, ErrNoDataSource
	}

	return d.wrapped.Read()
}

func isNilDataSource(source datasource.DataSource) bool {
	if source == nil {
		return true
	}

	value := reflect.ValueOf(source)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}
