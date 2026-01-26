package fs

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/creational/prototype/pkg/inode"
)

const defaultIndentStep = "  "

// Folder is a concrete prototype.
//
// Folder is recursive: it may contain both files and other folders.
type Folder struct {
	// name of the folder.
	name string

	// children contains folder contents (folders and files).
	children []inode.Inode
}

func NewFolder(name string, children ...inode.Inode) *Folder {
	childrenCopy := append([]inode.Inode(nil), children...)
	return &Folder{
		name:     name,
		children: childrenCopy,
	}
}

func (f *Folder) Print(w io.Writer, indentation string) {
	f.print(w, "", indentation)
}

func (f *Folder) print(w io.Writer, currentIndentation, indentStep string) {
	_, _ = fmt.Fprintln(w, currentIndentation+f.name)

	nextIndentation := currentIndentation + indentStep
	for _, child := range f.children {
		switch ch := child.(type) {
		case *Folder:
			ch.print(w, nextIndentation, indentStep)
		default:
			ch.Print(w, nextIndentation)
		}
	}
}

func (f *Folder) Clone() inode.Inode {
	clonedChildren := make([]inode.Inode, 0, len(f.children))
	for _, child := range f.children {
		clonedChildren = append(clonedChildren, child.Clone())
	}

	return &Folder{name: f.name + "_clone", children: clonedChildren}
}

func (f *Folder) String() string {
	var buf bytes.Buffer
	f.Print(&buf, defaultIndentStep)
	return strings.TrimSuffix(buf.String(), "\n")
}
