package fs

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/creational/prototype/pkg/inode"
)

// File is a concrete prototype.
//
// It implements inode.Inode and can be cloned without the caller knowing the concrete type.
type File struct {
	// name of the file.
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Print(w io.Writer, indentation string) {
	_, _ = fmt.Fprintln(w, indentation+f.name)
}

func (f *File) Clone() inode.Inode {
	return &File{name: f.name + "_clone"}
}

func (f *File) String() string {
	return f.name
}
