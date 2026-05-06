package filedata

import (
	"os"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/pkg/datasource"
)

type FileDataSource struct {
	path string
}

func NewFileDataSource(path string) datasource.DataSource {
	return &FileDataSource{path: path}
}

func (f *FileDataSource) Write(data []byte) error {
	return os.WriteFile(f.path, data, 0o644)
}

func (f *FileDataSource) Read() ([]byte, error) {
	return os.ReadFile(f.path)
}
